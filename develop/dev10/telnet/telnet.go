package telnet

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func CLI(args []string) int {
	var app App
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}

	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type App struct {
	timeout  time.Duration
	addresss string
}

func (app *App) fromArgs(args []string) error {
	fl := flag.NewFlagSet("telnet", flag.ContinueOnError)
	fl.DurationVar(&app.timeout, "timeout", time.Second*10, "timeout to connect to server")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	app.addresss = net.JoinHostPort(fl.Arg(0), fl.Arg(1))

	return nil
}

func (app *App) run() error {
	d := net.Dialer{
		Timeout: app.timeout,
	}

	conn, err := d.Dial("tcp", app.addresss)
	if err != nil {
		return err
	}

	fmt.Println("Connected to", conn.RemoteAddr())

	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	g := new(errgroup.Group)

	g.Go(func() error {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				fmt.Print("$: ")
				t, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						fmt.Println("\nEOF Interrupt signal")
						cancel()
						conn.Close()
						defer os.Exit(0)
					}
				}

				_, err = fmt.Fprint(conn, t)
				if err != nil {
					return err
				}
			}
		}
	})

	g.Go(func() error {
		reader := bufio.NewReader(conn)
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				t, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				fmt.Printf("Server: %s", t)
				fmt.Print("$: ")
			}
		}
	})

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		<-stop
		fmt.Print("\nInterrupt signal")
		cancel()
	}()

	err = g.Wait()
	if err != nil {
		if err == io.EOF {
			fmt.Println("Server closed.")
		}

		fmt.Println(err)
		return err
	}

	fmt.Println("\nDisconnected")

	return nil
}
