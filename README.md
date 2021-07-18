# Lamportのアルゴリズムを意識した全順序マルチキャスト通信
Lamportのアルゴリズムを意識してtickというタイムスタンプを用意した。これにより、マルチキャスト通信を用いて分散しているクライアント間で全順序性が保たれている。実行例を見ると、execute Taskが実行されるとき、実行順序が0->1で保たれていることが分かる。

## 実行方法
```bash
make build
make run
```

## 実行例
![lamport.png](./assets/lamport.png)

## 実行の流れ
![memo.png](./assets/memo.jpg)
