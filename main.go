package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	camera := rl.Camera3D{
		Position: rl.Vector3{
			X: 10,
			Y: 10,
			Z: 0,
		},
		Target: rl.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Up: rl.Vector3{
			X: 1,
			Y: 10,
			Z: 0,
		},
		Fovy:       90,
		Projection: 0,
	}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.DrawFPS(100, 100)
		rl.BeginMode3D(camera)
		rl.DrawGrid(10, 10)
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
