package main

import (
	"fmt"
	"os"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	// 读取WASM模块
	wasmBytes, err := os.ReadFile("./wasm/main.wasm")
	if err != nil {
		fmt.Println("Failed to read WASM file:", err)
		return
	}

	// 创建引擎和存储
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// 编译模块
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		fmt.Println("Failed to compile module:", err)
		return
	}

	// 创建WASI环境
	wasiEnv, err := wasmer.NewWasiStateBuilder("simple").Finalize()
	if err != nil {
		fmt.Println("Failed to create WASI environment:", err)
		return
	}

	// 创建导入对象
	importObject, err := wasiEnv.GenerateImportObject(store, module)
	if err != nil {
		fmt.Println("Failed to generate import object:", err)
		return
	}

	// 实例化模块
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		fmt.Println("Failed to instantiate module:", err)
		return
	}

	// 查找导出的函数
	sayHello, err := instance.Exports.GetFunction("_start")
	if err != nil {
		fmt.Println("Failed to find exported function:", err)
		return
	}

	// 调用导出的函数
	sayHello()

	// // 退出 WASI 环境
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// wasiEnv.Exit(ctx)
}