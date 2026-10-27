package main

import (
	"fmt"
	"math"
)

// ValidationError は図形の長さが不正だった場合のエラー情報を保持する構造体です。
type ValidationError struct {
	Message string  // エラーメッセージ
	Field   string  // 不正なフィールド（高さ・幅・半径など）
	Value   float64 // 不正な値（図形の長さ）
}

// error インターフェースを満たすための実装
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s (フィールド: %s, 値: %v)", e.Message, e.Field, e.Value)
}

// Shape は図形の面積と周長を計算するためのインターフェースです。
type Shape interface {
	Area() float64      // 面積を求める
	Perimeter() float64 // 周長を求める
}

// Rectangle は長方形を表す構造体です。
type Rectangle struct {
	Width  float64 // 幅
	Height float64 // 高さ
}

// Circle は円を表す構造体です。
type Circle struct {
	Radius float64 // 半径
}

func (p Rectangle) Area() float64 {
	return p.Width * p.Height
}

func (p Rectangle) Perimeter() float64 {
	return 2 * (p.Width + p.Height)
}

func (p Circle) Area() float64 {
	return math.Pi * p.Radius * p.Radius
}

func (p Circle) Perimeter() float64 {
	return 2 * math.Pi * p.Radius
}

// validateShape は図形のバリデーションを行う関数です。
func validateShape(s Shape) error {
	switch v := s.(type) {
	case Rectangle:
		if v.Width <= 0 {
			return &ValidationError{
				Message: "幅は0より大きい値を指定してください",
				Field:   "Width",
				Value:   v.Width,
			}
		}
		if v.Height <= 0 {
			return &ValidationError{
				Message: "高さは0より大きい値を指定してください",
				Field:   "Height",
				Value:   v.Height,
			}
		}

	case Circle:
		if v.Radius <= 0 {
			return &ValidationError{
				Message: "半径は0より大きい値を指定してください",
				Field:   "Radius",
				Value:   v.Radius,
			}
		}
	}

	return nil
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}

	for _, s := range shapes {
		if err := validateShape(s); err != nil {
			fmt.Println("バリデーションエラー:", err)
		} else {
			fmt.Printf("正常 - 面積: %.2f, 周長: %.2f\n", s.Area(), s.Perimeter())
		}
	}
}
