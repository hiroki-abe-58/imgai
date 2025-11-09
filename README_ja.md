# 🎨 imgai

AI搭載の画像処理CLIツール

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)

[English](README.md) | 日本語

Goで構築された、高速で効率的、かつユーザーフレンドリーなコマンドライン画像処理ツール。Apple Siliconに最適化されています。

## ✨ 機能

### 🖼️ 画像処理
- **リサイズ** - アスペクト比を維持または正確なサイズを指定
- **変換** - JPEG、PNG、WebP形式間の変換
- **品質制御** - 最適なファイルサイズのための圧縮調整

### 📊 バッチ処理
- **並列処理** - goroutineを活用した最大パフォーマンス
- **プログレスバー** - バッチ処理の視覚的フィードバック
- **Globパターン** - `*.jpg`パターンで複数ファイルを処理

### 🔒 プライバシーとメタデータ
- **EXIF読み取り** - カメラ設定、GPS、メタデータの表示
- **EXIF削除** - プライバシー保護のためすべてのメタデータを削除

### 🛡️ 安全機能
- **ドライランモード** - 実行前に操作をプレビュー
- **エラーハンドリング** - 詳細なエラーメッセージとリカバリー
- **自動命名** - スマートな出力ファイル名生成

## 🚀 インストール

### 前提条件
- Go 1.21以上
- macOS（Apple Silicon最適化）/ Linux / Windows

### ソースからビルド
```bash
git clone https://github.com/hiroki-abe-58/imgai.git
cd imgai
go build -o imgai
```

### クイックスタート
```bash
# 実行可能にしてPATHに移動
chmod +x imgai
sudo mv imgai /usr/local/bin/

# インストール確認
imgai --version
```

## 🌐 言語設定
```bash
# 日本語で使用
export IMGAI_LANG=ja
imgai resize photo.jpg --width 800

# 英語で使用（デフォルト）
export IMGAI_LANG=en
imgai resize photo.jpg --width 800
```

## 📖 使い方

### 画像をリサイズ
```bash
# 幅800pxにリサイズ（アスペクト比維持）
imgai resize photo.jpg --width 800

# 高さ600pxにリサイズ
imgai resize photo.jpg --height 600

# 正確なサイズにリサイズ
imgai resize photo.jpg --width 1920 --height 1080

# プログレスバー付きバッチリサイズ
imgai resize *.jpg --width 800

# リサイズ前にプレビュー
imgai resize *.jpg --width 800 --dry-run

# 8並列ワーカーで高速処理
imgai resize *.jpg --width 800 --workers 8
```

### フォーマット変換
```bash
# PNGに変換
imgai convert photo.jpg --format png

# カスタム品質でJPEGに変換
imgai convert image.png --format jpg --quality 85

# WebPに変換（モダンフォーマット）
imgai convert photo.jpg --format webp

# すべてのPNGをJPEGに一括変換
imgai convert *.png --format jpg --quality 90

# 変換前にプレビュー
imgai convert *.png --format jpg --dry-run
```

### メタデータ管理
```bash
# EXIFデータを表示
imgai exif photo.jpg

# すべてのメタデータを削除（プライバシーモード）
imgai strip photo.jpg

# メタデータを削除して新しいファイルに保存
imgai strip photo.jpg --output clean.jpg

# メタデータを一括削除
imgai strip *.jpg --workers 8

# メタデータ削除をプレビュー
imgai strip *.jpg --dry-run
```

## 🏗️ アーキテクチャ
```
imgai/
├── cmd/              # CLIコマンド
│   ├── root.go       # Cobraルートコマンド
│   ├── resize.go     # リサイズコマンド
│   ├── convert.go    # フォーマット変換
│   ├── exif.go       # EXIF読み取り
│   └── strip.go      # EXIF削除
├── pkg/              # コアパッケージ
│   ├── image/        # 画像処理ロジック
│   ├── batch/        # goroutineバッチ処理
│   ├── metadata/     # EXIF処理
│   └── i18n/         # 国際化対応
└── main.go           # エントリーポイント
```

## 🎯 主要設計原則

- **DRY** - Don't Repeat Yourself（繰り返しを避ける）
- **SOLID** - オブジェクト指向設計の5原則
- **関心の分離** - 明確なモジュール境界
- **エラーハンドリング** - 包括的なエラーメッセージ
- **クロスプラットフォーム** - macOS、Linux、Windowsで動作

## 🔧 開発

### ビルド
```bash
go build -o imgai
```

### テスト
```bash
go test ./...
```

## 📊 パフォーマンス

- **並列処理** - 複数ワーカーで最大8倍高速
- **メモリ効率** - ストリーミング画像処理
- **Apple Silicon最適化** - ネイティブARM64サポート
- **Goroutine** - バッチ処理の並行実行

## 📝 ライセンス

MIT License - 詳細は[LICENSE](LICENSE)を参照

## 🤝 コントリビューション

コントリビューションを歓迎します！Pull Requestをお気軽に送ってください。

1. リポジトリをフォーク
2. 機能ブランチを作成（`git checkout -b feature/amazing-feature`）
3. 変更をコミット（`git commit -m 'feat: add amazing feature'`）
4. ブランチにプッシュ（`git push origin feature/amazing-feature`）
5. Pull Requestを開く

## 👤 作者

**Hiroki Abe**
- GitHub: [@hiroki-abe-58](https://github.com/hiroki-abe-58)
- Repository: [imgai](https://github.com/hiroki-abe-58/imgai)

## 🙏 謝辞

以下のライブラリを使用して構築：
- [Cobra](https://github.com/spf13/cobra) - CLIフレームワーク
- [imaging](https://github.com/disintegration/imaging) - 画像処理
- [goexif](https://github.com/rwcarlsen/goexif) - EXIF処理
- [progressbar](https://github.com/schollz/progressbar) - プログレス表示

---

⭐ このプロジェクトが役に立つと思ったら、スターを付けてください！
