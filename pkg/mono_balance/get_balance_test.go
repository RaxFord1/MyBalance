package mono_balance_test

import (
	"MyBalance/internal/context"
	"MyBalance/pkg/core"
	"MyBalance/pkg/mono_balance"
	"testing"
)

func TestGetBalance(t *testing.T) {
	ctx := context.Named("test-GetBalance")
	ctx.SetString("mono_api", "fill_with_your_api_key_here")
	err := core.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}

	got, err := mono_balance.GetBalance(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(got)
}

//func TestGetBalance_Concurrently(t *testing.T) {
//	ctx := context.Named("test-GetBalance-concurrently")
//  // do not test it, or get banned
//	for i := 0; i < 100; i++ {
//		go func() {
//			_, err := GetBalance(ctx)
//			if err != nil {
//
//			}
//		}()
//	}
//
//	got, err := GetBalance(ctx)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Log(got)
//}
