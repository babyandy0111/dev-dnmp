# 啟動方式

1. `docker-compose up -d --build` 這會將同目錄的 `plugins` 編譯進 WordPress image 內，並透過 docker-compose 啟動
2. 打開瀏覽器，輸入 `http://127.0.0.1:8003` 可見到 WordPress 安裝畫面
3. 選擇安裝英文版本（為了避免 WPGraphQL 在安裝後啟動不能的問題，可以在安裝後切換語言介面為中文）
3. 按照以下資訊輸入安裝資訊：
    - 資料庫名稱： `wordpress`
    - 資料庫帳賀： `root`
    - 資料庫密碼： `12345678`
    - 資料庫主機： `mysql`
    - 資料表前綴： `wp_`
4. 輸入網站資訊：按照需要自己填入就好
5. 設定完成，登入到管理介面
6. 到 `設定 > 永久連結` 將預設的 `/?p=123` 換成其他方式，才能確保 /graphql endpoint 可以動作
7. 啟用 plugin：`Add WPGraphQL SEO`、`User Role Editor`、`WP GraphQL`、`Yoast SEO`
8. 依照實際需求到各 plugin 更改設定。