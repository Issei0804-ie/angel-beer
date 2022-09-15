# angel-beer

### agentのインストール方法

```
go install github.com/Issei0804-ie/angel-beer/agent@latest 
```

### supervisorのインストール方法
```
go install github.com/Issei0804-ie/angel-beer/supervisor@latest 
```

# 開発時初期設定

## git commit templatesの有効化

.gitmessageに自分の行を追加後、以下を実行。

```
$git config --local commit.template .gitmessage
```