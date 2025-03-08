package parser

import (
	"log"
	"net/http"
	"online-Market_project_Golang-Backent/internal/db"
	"online-Market_project_Golang-Backent/internal/models"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseLenta парсит данные о товарах с указанной категории на lenta.com
func ParseLenta(categoryURL string) {
	// Отправляем HTTP-запрос к странице категории
	resp, err := http.Get(categoryURL)
	if err != nil {
		log.Printf("Ошибка при запросе к %s: %v", categoryURL, err)
		return
	}
	defer resp.Body.Close()

	// Парсим HTML с помощью goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Ошибка при разборе HTML: %v", err)
		return
	}

	// Ищем элементы товаров на странице (классы взяты с сайта lenta.com)
	doc.Find("div.sku-card").Each(func(i int, s *goquery.Selection) {
		// Извлекаем название
		name := strings.TrimSpace(s.Find("a.sku-card__title").Text())
		if name == "" {
			return // Пропускаем, если название пустое
		}

		// Извлекаем цену
		priceStr := strings.TrimSpace(s.Find("div.sku-card__price-primary").Text())
		price := parsePrice(priceStr)

		// Извлекаем ссылку
		link, exists := s.Find("a.sku-card__title").Attr("href")
		if !exists {
			link = ""
		}
		fullLink := "https://lenta.com" + link

		// Создаём объект Product
		product := models.Product{
			Name:          name,
			Price:         price,
			Description:   "Товар с категории: " + categoryURL,
			AgeRestricted: false, // Можно добавить логику для 18+
		}

		// Сохраняем в базу данных
		if err := db.DB.Create(&product).Error; err != nil {
			log.Printf("Ошибка при сохранении товара %s: %v", name, err)
		} else {
			log.Printf("Успешно сохранён товар: %s, Цена: %.2f, Ссылка: %s", name, price, fullLink)
		}
	})
}

// parsePrice преобразует строку цены в float64
func parsePrice(priceStr string) float64 {
	// Удаляем лишние символы (₽, пробелы) и преобразуем в число
	priceStr = strings.ReplaceAll(priceStr, " ₽", "")
	priceStr = strings.ReplaceAll(priceStr, " ", "")
	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Printf("Ошибка преобразования цены '%s': %v", priceStr, err)
		return 0
	}
	return price
}
