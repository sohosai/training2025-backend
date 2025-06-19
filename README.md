# IPC / jsys 合同 Web 研修 2025 掘っ立てバックエンド

2025/6/21, 22に開催される IPC / jsys 合同 Web 研修 2025 のための掘っ立てバックエンドのリポジトリです。

超簡易的SNSサービスを提供します。

`training25api.playground.sohosai.com:8080`あたりでホストされる予定

## エンドポイント・操作一覧

- `/posts`
  - `GET` (HTTPメソッド)
    - 投稿の一覧を取得できる
    - リクエスト: ボディなし
    - レスポンス
      ```http
      [
        {
          "id": "投稿ID(UUID) (string)",
          "text": "投稿内容の文字列 (string)",
          "created_by": "投稿をした人(特にユーザ認証などはないので適当にユーザ名でも入れてください) (string)",
          "created_at": "投稿時刻 (string)",
          "updated_at": "最終更新時刻 (string)",
        },
        {
          "id": "2個目の投稿のID",
          "text": "2個目の投稿内容",
          "created_by": "2個目の投稿をした人",
          "created_at": "2個目の投稿時刻",
          "updated_at": "2個目の最終更新時刻",
        },
        {
          "id": "3個目の投稿のID",
          "text": "3個目の投稿内容",
          "created_by": "3個目の投稿をした人",
          "created_at": "3個目の投稿時刻",
          "updated_at": "3個目の最終更新時刻",
        },
      ]
      ```
  - `POST` (HTTPメソッド) **実装予定**
    - 投稿できる
    - リクエスト
      ```http
      {
        "text": "投稿内容の文字列 (string)",
        "created_by": "投稿をした人(特にユーザ認証などはないので適当にユーザ名でも入れてください) (string)",
      }
      ```
    - レスポンス
      ```http
      {
        "id": "割り当てられた投稿のID(UUID) (string)",
      }
      ```
- `/posts/投稿のID`
  - `GET` (HTTPメソッド) **実装するかも**
    - IDで指定した投稿を取得できる
    - リクエスト: ボディなし
    - レスポンス
      ```http
      {
        "id": "投稿ID(UUID) (string)",
        "text": "投稿内容の文字列 (string)",
        "created_by": "投稿をした人(特にユーザ認証などはないので適当にユーザ名でも入れてください) (string)",
        "created_at": "投稿時刻 (string)",
        "updated_at": "最終更新時刻 (string)",
      },
      ```
  - `POST` (HTTPメソッド) **実装するかも**
    - IDで指定した投稿を更新できる
    - リクエスト
      ```http
      {
        "text": "投稿内容の文字列 (string)",
      }
      ```
    - レスポンス: ボディなし
