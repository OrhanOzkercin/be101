# Blog Yazısı İçin Gerekli Alanlar ve API Tanımları

## **Temel Alanlar (Minimum Gereksinimler)**

1. **ID** → (UUID veya integer) → Blog yazısının benzersiz kimliği
2. **Title** → (string max 255 karakter) → Blog başlığı
3. **Content** → (text max 10000 karakter) → Blog içeriği
4. **AuthorID** → (integer/UUID) → Yazıyı oluşturan kullanıcının kimliği
5. **CreatedAt** → (timestamp) → Oluşturulma tarihi
6. **UpdatedAt** → (timestamp) → Son güncellenme tarihi

---

## **Gelişmiş Alanlar (İlerleyen Aşamalarda Eklenebilir)**

7. **Slug** → (string) → SEO uyumlu URL için
8. **Excerpt (Özet)** → (string) → İçeriğin kısa özeti
9. **CoverImage** → (string) → Kapak görseli URL’si
10. **Tags** → (array/string) → Bloga ait etiketler (örn. ["Golang", "Backend"])
11. **Category** → (string) → Blogun kategorisi
12. **Views** → (integer) → Görüntüleme sayısı
13. **Likes** → (integer) → Beğeni sayısı
14. **CommentsCount** → (integer) → Yorum sayısı
15. **IsPublished** → (boolean) → Yayınlanmış mı, draft mı?
16. **PublishedAt** → (timestamp) → Yayınlanma tarihi

---

## **Önerilen İlk API'lar**

### **1. Blog Kaydetme (Create Post)**

- **Endpoint:** `POST /posts`
- **Gereksinimler:** `title`, `content`, `authorID`
- **Yanıt:** Kaydedilen blogun bilgileri

### **2. Tüm Blogları Getirme (Get All Posts)**

- **Endpoint:** `GET /posts`
- **Opsiyonel Filtreler:** `authorID`, `category`, `tag`, `isPublished`
- **Sayfalama:** `limit`, `offset` veya `cursor-based pagination`

Bu iki API tamamlandıktan sonra yorum sistemi, yetkilendirme ve diğer gelişmiş özellikler eklenebilir. 🚀
