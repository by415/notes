package solutions

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
)

// 验证context WithCancel 方法
func TestWithCancel(t *testing.T) {
	funcWork := func(ctx context.Context) {
		defer func(ctx context.Context) {
			fmt.Println("defer")
			time.Sleep(time.Second * 2)
			fmt.Println("WithCancel = ", ctx.Err())
		}(ctx)
		for true {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("workFunc")
			}
			time.Sleep(time.Second * 1)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	go funcWork(ctx)
	time.Sleep(time.Second * 2)
	cancel()
}

func TestWithTimeout(t *testing.T) {
	funcWork := func(ctx context.Context) {
		defer func(ctx context.Context) {
			fmt.Println("defer")
			time.Sleep(time.Second * 2)
			fmt.Println("WithTimeout = ", ctx.Err())
		}(ctx)
		for true {
			select {
			case <-ctx.Done():

				return
			default:
				fmt.Println("workFunc")
			}
			time.Sleep(time.Second * 1)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	go funcWork(ctx)
	time.Sleep(time.Second * 4)
	fmt.Println("endl")
	cancel()
}

func TestWithDeadline(t *testing.T) {
	funcWork := func(ctx context.Context) {
		defer func(ctx context.Context) {
			fmt.Println("defer")
			time.Sleep(time.Second * 2)
			fmt.Println("WithDeadline = ", ctx.Err())
		}(ctx)
		for true {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("workFunc")
			}
			time.Sleep(time.Second * 1)
		}
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	go funcWork(ctx)
	time.Sleep(time.Second * 4)
	fmt.Println("endl")
	cancel()
}

// 参数使用
// context value的作用是向下传递参数的吧。这样使用不合理
func process(ctx *context.Context) {
	for true {
		time.Sleep(time.Second)
		ret, ok := (*ctx).Value("trace_id").(int)
		if !ok {
			ret = ret + 1
		}

		select {
		case <-(*ctx).Done():
			return
		default:
		}
		ret = ret + 1

		fmt.Printf("ret:%d\n", ret)
		*ctx = context.WithValue((*ctx), "trace_id", 3)

		s, _ := (*ctx).Value("session").(string)
		fmt.Printf("session:%s\n", s)
	}
}

// 验证context value
func TestWithValue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	ctx = context.WithValue(ctx, "trace_id", 2222222)
	ctx = context.WithValue(ctx, "session", "sdlkfjkaslfsalfsafjalskfj")

	go process(&ctx)
	time.Sleep(4 * time.Second)
	fmt.Println("xxxxxxxxxxxxxxxxxx", ctx.Value("trace_id"))
	cancel()

	defer func(ctx context.Context) {
		fmt.Println("xxxxxxxxxxxxxxxxxx", ctx.Value("trace_id"))
		fmt.Println(ctx.Value("session"), ctx.Err())
	}(ctx)
}

// 利用context监控调用堆栈
func TestValue2(t *testing.T) {
	// context实例
	ctx := context.Background()
	go test01(AppendContextValue(ctx, "func", runFuncName()))

	go test02(AppendContextValue(ctx, "func", runFuncName()))

	time.Sleep(2 * time.Second)
	fmt.Println("in main")
	PrintCtxValue(ctx)
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// 打印 value
func PrintCtxValue(ctx context.Context) {
	x := ctx.Value("func")
	if x != nil {
		fmt.Println(x)
	}
}

// 更新value值
func AppendContextValue(ctx context.Context, key string, value string) context.Context {
	xx := ctx.Value(key)
	strValue := ""

	if s, ok := xx.(string); ok {
		strValue = s
	}
	strValue = strValue + "\n" + value
	ctx = context.WithValue(ctx, key, strValue)
	return ctx
}

func test01(ctx context.Context) {
	go test02(AppendContextValue(ctx, "func", runFuncName()))
}

func test02(ctx context.Context) {
	go test05(AppendContextValue(ctx, "func", runFuncName()))
}

func test05(ctx context.Context) {
	test06 := func(ctx context.Context) {
		fmt.Println("inner funcName = ", runFuncName())
		PrintCtxValue(AppendContextValue(ctx, "func", runFuncName()))
		fmt.Printf("%s",debug.Stack())
	}
	go test06(AppendContextValue(ctx, "func", runFuncName()))
}

// 验证withvalue的线程安全问题
func TestWithValue3(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "1111")
	go func(ctx1 context.Context) {
		i := 1
		for {
			time.Sleep(1 * time.Second)
			ctx = context.WithValue(ctx1, "test", i)
			i++
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(ctx.Value("test"))
		}
	}(ctx)
	time.Sleep(20 * time.Second)
}
