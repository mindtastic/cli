package main

import (
	"context"
	"fmt"
	client "github.com/mindtastic/cli/client"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type nf struct {
	f    func(ctx context.Context) error
	name string
}

type BenchmarkingClient struct {
	c     *client.APIClient
	token string
	l     *log.Logger
	calls []nf
}

func NewClient(id int, out io.Writer, client *client.APIClient) *BenchmarkingClient {
	ret := &BenchmarkingClient{
		c: client,
		l: log.New(out, fmt.Sprintf("[%d] ", id), log.Ltime|log.Lmicroseconds|log.LUTC),
	}
	ret.calls = []nf{{f: ret.mood, name: "mood"}, {f: ret.users, name: "user"}, {f: ret.wiki, name: "wiki"}, {f: ret.motivation, name: "moti"}}
	return ret
}

func (b *BenchmarkingClient) Init(ctx context.Context) {
	b.register(ctx)
}

func (b *BenchmarkingClient) Run(ctx context.Context) {
	if b.token != "" {
		return
	}
	ctx = context.WithValue(ctx, client.ContextAccessToken, b.token)
	for {
		if err := ctx.Err(); err != nil {
			b.l.Printf("shutting down: %v", err)
			return
		}

		b.run(ctx)
		time.Sleep(time.Second)
	}
}

func (b *BenchmarkingClient) run(ctx context.Context) {
	r := rand.Intn(len(b.calls))
	ts := time.Now()
	nf := b.calls[r]
	err := nf.f(ctx)
	te := time.Now()
	td := te.Sub(ts)
	b.l.Printf("| %s | %d | %v", nf.name, td.Milliseconds(), err)
}

func (b *BenchmarkingClient) register(ctx context.Context) {
	req := b.c.AuthenticationApi.BouncerInitRegistrationApi(ctx)
	flow, res, err := req.Execute()
	if err != nil {
		b.l.Printf("error initializing registration: %v", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		b.l.Printf("error initializing registration: %s", res.Status)
		return
	}
	reg := b.c.AuthenticationApi.BouncerRegister(ctx)
	reg = reg.Flow(flow.Id)
	auth, res, err := reg.Execute()
	if err != nil {
		b.l.Printf("error executing registration request: %v", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		b.l.Printf("error executing registration request: %s", res.Status)
		return
	}
	b.token = *auth.SessionToken
}

func (b *BenchmarkingClient) mood(ctx context.Context) error {
	req := b.c.MoodDiaryServiceApi.DiaryMoodGetMany(ctx)
	_, res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("error executing mood request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error executing mood request: %s", res.Status)
	}
	return nil
}

func (b *BenchmarkingClient) users(ctx context.Context) error {
	req := b.c.UserServiceApi.UserDataGet(ctx)
	_, res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("error executing users request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error executing users request: %s", res.Status)
	}
	return nil
}

func (b *BenchmarkingClient) wiki(ctx context.Context) error {
	req := b.c.WikiServiceApi.WikiList(ctx)
	_, res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("error executing wiki request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error executing wiki request: %s", res.Status)
	}
	return nil
}

func (b *BenchmarkingClient) motivation(ctx context.Context) error {
	req := b.c.MotivatorServiceApi.MotivatorGet(ctx)
	_, res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("error executing motivator request: %v", err)

	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error executing motivator request: %s", res.Status)
	}
	return nil
}
