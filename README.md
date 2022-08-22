# godoc-sandbox

godoc を使い方を試すためのプロジェクトです。

## godoc
Mac/Linux 想定で説明しています。

Windows の場合は適宜読み替えてください。

### インストール
godoc コマンドのインストール
```
go install golang.org/x/tools/cmd/godoc@latest
```

### プロジェクトを反映
GOROOT の src ディレクトリ配下に反映

プロジェクトディレクトリ内がカレントディレクトリとしています。
```
ln -s $(PWD) ${GOROOT}/src/
```

### 起動
あとは godoc コマンドを実行するだけです。
```
godoc
```

ブラウザにアクセスし、`http://localhost:<port>/pkg/<project 名>` にアクセスしてください。

port はデフォルトでは 6060, project 名は今回 `godoc-sandbox` です。

ポート指定などもできます。詳しくはヘルプ（godoc --help）を参照。
