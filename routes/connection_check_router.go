package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func APIConnectionCheck(api fiber.Router) {
    api.Get("/connection-test/soccerway", func(c *fiber.Ctx) error {

        client := &http.Client{}

        req, err := http.NewRequest("GET", "https://ru.soccerway.com/a/block_h2h_matches?block_id=page_match_1_block_h2hsection_head2head_7_block_h2h_matches_1&action=changePage&callback_params=%7B%22page%22%3A+-1%2C+%22block_service_id%22%3A+%22match_h2h_comparison_block_h2hmatches%22%2C+%22team_A_id%22%3A+43000%2C+%22team_B_id%22%3A+43009%7D&params=%7B%22page%22%3A+0%7D", nil)
        if err != nil {
                log.Error(err)
        }

        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/131.0")

        resp, err := client.Do(req)
        if err != nil {
                log.Error(err)
        }

        defer resp.Body.Close()


        return c.JSON(fiber.Map{
            "status": resp.StatusCode,
        })
    })

    api.Get("/connection-test/marafon", func(c *fiber.Ctx) error {

        resp, err := http.Get("https://www.marathonbet.ru/su/betting/Football+-+11?page=0&pageAction=getPage")
        if err != nil {
            log.Error("Error requesting to marafon!")
        }

        return c.JSON(fiber.Map{
            "status": resp.StatusCode,
        })
    })
}
