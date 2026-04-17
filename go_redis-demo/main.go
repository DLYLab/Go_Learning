package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// 定义一个背景上下文
var ctx = context.Background()

func main() {
	// 1. 初始化连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 密码，没有则留空
		DB:       0,                // 默认数据库
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接 Redis 失败: %v\n", err)
		return
	}
	fmt.Println("成功连接到 Redis!")

	// --- 基础操作演示 ---

	// 2. String (字符串) - 最简单的键值对
	fmt.Println("\n[String 操作]")
	err = rdb.Set(ctx, "username", "tony_stark", 10*time.Second).Err() // 设置 10 秒过期
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "username").Result()
	fmt.Printf("Get username: %s\n", val)

	// 3. Hash (哈希) - 适合存储对象
	fmt.Println("\n[Hash 操作]")
	rdb.HSet(ctx, "user:100", "name", "Iron Man", "age", 40)
	userFields, _ := rdb.HGetAll(ctx, "user:100").Result()
	fmt.Printf("User:100 详情: %v\n", userFields)

	// 4. List (列表) - 简单的队列或栈
	fmt.Println("\n[List 操作]")
	rdb.RPush(ctx, "tasks", "task1", "task2", "task3") // 从右侧插入
	task, _ := rdb.LPop(ctx, "tasks").Result()         // 从左侧弹出
	fmt.Printf("弹出任务: %s\n", task)
	
	allTasks, _ := rdb.LRange(ctx, "tasks", 0, -1).Result()
	fmt.Printf("剩余任务列表: %v\n", allTasks)

	// 5. Set (集合) - 无序且唯一
	fmt.Println("\n[Set 操作]")
	rdb.SAdd(ctx, "tags", "tech", "golang", "redis", "tech") // 重复的 tech 不会被加入
	tags, _ := rdb.SMembers(ctx, "tags").Result()
	fmt.Printf("所有标签: %v\n", tags)
	
	isMember, _ := rdb.SIsMember(ctx, "tags", "golang").Result()
	fmt.Printf("包含 golang 标签吗? %v\n", isMember)

	// 清理数据（可选）
	// rdb.Del(ctx, "username", "user:100", "tasks", "tags")
}