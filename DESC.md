# 説明

## 工夫点

- フレームワークとしてechoを採用した。
- ORMとしてgormを採用した。
- config・環境変数の管理にgodotenv, envconfigを採用した。これは、開発をDockerで行ったため、環境の切り替えを環境変数で行ったためである。
- コアロジックとメインの起点を分けることで、コードの再利用性を高めることを意識した。
- 今回は時間が足りなかったが、godoc形式のコメントを書くことを意識した。
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)や、[Qiitaの記事](https://qiita.com/rema424/items/9ffbdf584b705cae6a19#3-%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E7%AE%A1%E7%90%86)を参考にしたディレクトリ構成とした。
