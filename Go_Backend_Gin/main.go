/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-27 00:30:50  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-27 00:30:50  */ /* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-27 00:30:49  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-27 00:30:49  */ /* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-27 00:30:46  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-27 00:30:46  */
package main

import (
	"app/config"
	"app/route"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	config.LoadEnvironmentVar()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(gin.Logger())

	route.Router(r)

	log.Println("Server is started in PORT 8080 ...👨‍💻 ")
	if e := r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT")); e != nil {
		log.Fatalln("🔥❌ ERROR   👨‍💻 : ", e)
	}

}
