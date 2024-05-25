package mono_balance_test

import (
	"MyBalance/internal/core"
	"MyBalance/internal/http/context"
	"MyBalance/internal/projkeys"
	"MyBalance/pkg/mono/mono_balance"
	"fmt"
	"testing"
	"time"
)

func TestGetBalance(t *testing.T) {
	ctx := context.Named("test-GetBalance")

	err := core.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ctx.SetString(projkeys.ClientID, "clientId")

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

func Test_formatCardInfo(t *testing.T) {
	fmt.Println(fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04")))
}
