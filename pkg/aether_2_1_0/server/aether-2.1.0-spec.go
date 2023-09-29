// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xdC3PjuJH+KyjdVa3nIkq2Z7S169RW4rG9c67YM44fd8mNXVOwCMnI8rUg5Ec8/u9X",
	"APgASPApkH4sU5UdWSLRQKP760aj0XgczX038D3k0XC08zgKIIEuoojwv5BHEQkIDpGFbfbF3Pco8ij7",
	"SNE9nQYOxN6fwfwGkhDRX1Z0Yf3EfgznN8iF/LGHAI12RiEl2FuOnp6exiMbhXOCA4p9b7Sj0gAbFJIl",
	"ogB7wPf80Jr73gIv343GI8weDiC9GY1HHnRR9tXReETQ7ytMkD3aoWSFGCn2DQrpR9/GiI/oNPni4dtu",
	"EDh4DkU/lLHB9Jfpv0Lxczqk/yRoMdoZ/cc0Zd1U/BpO5Taf+GgLKH478OzAx4KeadJp4/X68O34mnTa",
	"D04g35czTJExurwxPY1ve77noTnFt5g+WCEit3hulrCeQoPefNvzCbI+LLvvVUKpce9mvfVu1qx3p9Cz",
	"Zsv+JjZLsKCv+8h8d6I2SylaS+Kvgi7oRi3XoN7h4FUCdfpiEt70rRf04tANsWWjBfawUTOjbbyoD4Fl",
	"+y7EpqknzRbQPfY9TH1u9o0SltqtpPztwF4iy+5AFAtoFPToDLvWHBLbbB+SVouoOsaHLZoso9chBmna",
	"L+3Jr9ihiHTRh6jlUurGQSdttpTuCcE+YcaKErhY4LlFVk4XUlBAqHnfvn3qhlPFxFr08bjPPlbO8T9g",
	"0Ilu8XaLKLvQcaw5chzDlNN2CyhfBAuzJFmDeVrnyA0caHBFkjRYTMuoXCmNamhGIjZ3YBiaI6q0+iTI",
	"Rj+yd3dtm7sn0DkhfoAIfThIl+2H3CKqAYEvgXjceQBhgOZ48QAgUKMEPr1BBNAb6AF6g4CNFnDlULDh",
	"e84D8D1wsnu+99/ARfTGt9+NxqNAUI5iALmohkofesC/RYRgGwF/wQloYxSsYYqpo4tEZKMe6ZNV7Eje",
	"9a//heZ09DTWvHLhzW+gt0Sa3p/7wE8Z6LuYgh/iuMgPIGXEGEDbZoNzAfXBD6u4xR+Ag0Oa49mqmOIu",
	"mPuuC0GIAkQgRTZvgbEueQm40LMh9ckDgJQSfL2iCHjQRaHEw5REM/5d5N+TeKfGeGD27Wh4pfGMEoq5",
	"aNZRNHRJn0KwEWJv6aC8JELbJkjoYoan4gchfziUm5uAXQ8cnoDoXXCH6Q0IV9ceomDj8OT2A/AJODy5",
	"/VGWz5hQjrNjWfO16nC4DxY+yXdjNB658P4IeUt6M9r58f145GIv/nNrPAogpYiwFr5C699XG+y/m9bP",
	"1tVfok9X7/5L7qHajWxHx6N7a+lb0ZeMy39DD/zbB+gt4+8zjWRnJzs0x/eWQPoKLDBybHVgW5vbH7JD",
	"i/sst6brsNw1+VnWLxwGDnywRCQz27HoV64iTD1XIQLYA58uDtnk7h0dqn38abOwhzIZzeQjKQjZOKbH",
	"ZmEUR1lF3PWrJGmZybiSdFhSyqzOMr4x/LDclUNxwJ7nwVxFlZXgaec6zeGM6UDMrWKFHsStlrjpgcZO",
	"4C4RS7MYI5NvCTCIeR4WtpFH8QKzJeB45AoXrl1MfDwKfEIt5Gk4cgRDCtjPbDYIk07BEOyu3NHOj7PZ",
	"+xnniPh7Mx1o0mQySuxRtEQpvZBCQvMUf8VkHZKiVS1R4lN/7juCJHfXmLOydzLKattJ9GSFMKiy8F6W",
	"hfO9k+8X+yfSxCfUq9QmeZB5HFl3OWvpvd+YskSPAfFYSlN9XUfYQXAhZGZK5zvK87kvhMgm71gEhb5z",
	"iwh7WQjl9HZ7sjW5nz4qvujTVPbsH1VMfprGvGVvJbrxNFX7cgudFQpzSK9qkzT/eqRPMbsm5LOv3VVI",
	"RztfH0cHhPjkGIUhXHIA+yx2AFOqgD0KrhFwmFPElwYMsX9fQYfNUqIRT+P4XW/lXiOyMWEc2kkbegcu",
	"V5ub7+e/gPwDyLPfjZ6uCoyRMIh52M+bD/YV39OkyA3b7euV81h0JWE0JAQyMFt5+PcVOhRUYx7/hh7C",
	"7HxeNTPD8SpWHfixgA1wjSlbF+SNpX/nOdj7Lf/mCSKWCPOBABHZ7QTHH09B/CKwIYWAtc3Q6jpgarbw",
	"iQupAJ8fP4z0YJVQ1kHVKqjTKWs30ynxWusuRVTzHaqaacb68oVPgVzqlyptRbNIIlsIYt5xLJLFIwQX",
	"p2jxpcDhEl4acOA1cgSKxQv6CEajpXJOMvkbGk1mXxc2EY9evKzxfngP8q2WdyxuVbysXRlnp17hSlg8",
	"8Rm+qKTDujKgzoEkBZluZGTgaTyKd//NOPBqGKVwYR5iior993lBnkDL7f8XvP5EtYcmMC9a7cVvpts7",
	"zfZKk1ZexoIE5/dnm2+7jkdY3mKtv3OacMNVtkob7YCOR6G0t1h7yzAhzfShMu7DHjId8IkJt1uIxW+z",
	"ETh1JVnscSQjV/YzGmxTJC2sxL5Ene0GfbAkHoXkMXNQNOghQyAUFnD9A3PoeT73k1chsplSuT5Bkb/s",
	"IcCZyb6GfNIld3nurzy6wb7c4Q99DRG15r5HIfbCjYXvOP4d9pZWiK8d7C13dtJHp/yjDBzTv+a+GoM6",
	"T727ir7kDVvYfgd+AZvCJy/N9FKFW36KL+JuEHD9a+wgMPcJ0tiFJDlrzZyrsWhrtjSQITUekVzik6F8",
	"JlUg9VzVuB610tsyGxY3CHxYcrbLS/1MoHw+twLiu2yluwqtFdF4ZxenR8zA7+7tgfRROcKcb6PKaK4I",
	"zkZJVZqxFw4uTo80gaZ6BOowO+Ffc6bPCpg+G5hei+mzhkzPJyNWY8/9bhDk+H8Pg8AqZgN7BxTwQn3V",
	"HEMyYyviS5r7p3fCBaKHrzCOLkxRlb8kHjPtMaXE2/lM6fsvyQFHWMPL44PDLDcz8/ujSvGDwk3Guset",
	"D+OtH58k/nFaVRPOH8p41frAb4hd/oDs0kbvlEd7J5PpZDKNH+YfWkd3mS80fYz8yKepYNX0MZnqp5RO",
	"HMYdj0IaZbnUXfudUZHFonqvqTxl/NdY/5u7sVcqgqSrzN427mVv+dUCVNtkau0ivybaCZZ1hHlJR9ZC",
	"vqSVF4R/cuRAjzPJIyBElGJvybc3KLARRcTFHuIeROCLLbM0MwSF40sPhuAOOQ6AociUEg3F+CK1yMco",
	"g2XSsTpgljydfjKMadGK8DEzlU8S6RTgauzIFqTxv5ltv2ZMrLXhp9FFZbB6KxBheP08j7KjJJqdNcmd",
	"xJ4MSOLlMggv9OXiSU48j3VcsKyuTCaxmZZb636ONb5B8lHSHeTBawcp+/TCSGcXYeyxjIOWLsV4GwmP",
	"rn3fQdBr7EGo87++CJXt0FbKUaOdseLjStUDbbNnpnCyPlvKdwYzvtB6DCgfefshJ3DUaODRhnXnbuV6",
	"e+DRBDzvnrfUiQ73uBVYSLsavVIKDwU74LnNq+pwRFsJL5DtPhQ5WpvlxkYQtAHPh08yrsN0rEUbn7r8",
	"buY3Jj8DtnZchWDjnyj8/tn/fmhHpjVx66Uf+JF8aH/xnId4oDGXUnp6p7gwM5r1BgdJ9jMMQ7z0kB0H",
	"0pLhFRGW2tZQdmDINzWKeHF+gwDFLhf8i/M94YindMEdDAFrA8jjS3lzeWk/fniy2D/b8T/n4p8d5Z+N",
	"y8vJ5aX9p3d/2fi/718vL/90eWldKU+8KxljZhRl6fR5SSrSJM1R1s7RM9peQoQvNtPNWqjN2ki9o/zE",
	"SZ6TOFSC50DK4lQB7P22nPT4888/6+FMoqfD1LjFfHDrjAe3XEijAy4YkXEUL7hGIbjx70TEIASQIBDA",
	"+W98q476AHqAvT4Bx1G2G7qHc+o8gK0ZL8IB5xSREDi+t5yAX33CfncDB43B3t7e58+fDw4Ozvj/Muv0",
	"rZm6gJ2lo4yGoVEWdz7Pjy7ZRVt5lDyAuW+jUT42916Oy7F2qsIc7BlG0ism6SF655Pfikhuj1WiXg2i",
	"3lyT9ShNesIbMQT2fNZSZdWmUL3kQEBv0S4cRCGB0kMqLvYsgftq+u7B592PRwe5DF7+Ag6Zw4NvUWQx",
	"8uHbdIJEO9+Pdw8/nx983v28d/B9//Asajs9wiJ1o2rmlIeHYyBVUSjb07Bl//Pn6e7J50za9XYhDa/G",
	"6D0xai+0AoJdSB7yZKMfgO2FIETkFhEQdzuhlL6uHUxohWjue7a2/eSnCgppE2Vhu8rgaKJhpiOjmUBb",
	"m7Co0gRDV7pSFHzrw8/bWeWOzCKgBHqhi8OQKczKw7TmOQFGQ2ctxQk2zXSJk23M9jmOP+eLjiANcYIF",
	"8V0F7Tc2ON5//7pl/XwlPm7xf8Tn7a+b1of48+zrpjW7end5OXn3+P6p+YvTiNi77xtft6ztq/iP9183",
	"re2rd/IJvGh8ldtPwe0HKyBoge/zyyKuYZl5j9rNmZ18+LZe1CSTnFa4ZpItR/Nlk1Q8pKjjbRZPCm+q",
	"1k/l4xN5oo1HlhtPi1FIWWGlA1ALrJR5zlJ+Yc5ptpfImjurkCJSN9uDvQOid/RpH2Wt1s3/UCu3tCzI",
	"kuxo8eZSPtQdavpGyUCLm22VeXEsz5beWyyucKOXZj5j0lhebwLGCznImDK95CyjnuumDzWqXTF2rjG7",
	"7lHJXBXKrCKQDY1PmQ43l+3mZqmoqFOtwbY6c5Xjail/zkqSUmKz5QL2xKDa7c8oN97/YtOyB4k9ST5F",
	"W/SQIGDjkD1pgzvsOCBJvuZvYw9A7yEOHkZ/8czmSZ3dtPEIz+da+NnbU9P3hVCoLNrOLOV+1qRQ/TTe",
	"3npin3etX6/kxQenW+nM8qeiYxbFgbCyTmYDUwV5XjM1zyusk+cVJnle1ccfRNfMH4Fw1zkBIbCAr91y",
	"5wpcHUif5fLTmidmqQdI6uBQcxhO69oVDKCdV+3Wgtj4NElvMThOMWx+Jk2zM/EQ8JN88hsgegN8WE5n",
	"yxg5edEYehOd9xgx1GPr9K8jnmE+W8qiMy/N/ZfiLgIurWt0A2+xrzkQvB/VaYqfAHgBPB8seDE9QFYO",
	"CoEL6fwmA1SbsmbtH3z+p7V7dPR99+joy/9Kn04uPh4d7inJGpkOVWFC8uCLP7jX4PhdvmjiSzuEt0jK",
	"NDaqvpienrsmTesnMh0rqJK4VuG+5GyZxqLw94CNFwtEmJsNqU8m4NB1V5TZ96xxKbEtwh5fXk6uHrfG",
	"SopzWG2Uw/TQXLXhczpIY09It7R7TprEHupMEOfz9CyCPfZWMZe363J5W+FyWB3IC5Ujgvr0slWICAgc",
	"yFM4IQXYDRzkMokT7L+dL+T6aMGiVi7mKliw/xtLKxNH+B5jvj9xAmm+2H1UirJZhclsHlgOqbkkS6LC",
	"+JnzZ5wWa0x9kdpsSTtHlwI1AZ9EKhR35XlyRQAJxXMciGSYS4/ZVKHoOC4dyD19biAAUykRzD4nK1SZ",
	"mljQOzU7EST5lZqUaRNpilFCYT77swvZKs5hzPxmIGExx7mGaYvida1QrpX7WmSxn01MW3jwmkrNlXxa",
	"Nwmwnm8vVYUu4qZSESXaeboD0LOnPgE28h4m4JQ7q2x9j+7RfEV5gsalF3sVwCc2IhPAvPIFr6nF3APW",
	"FvdvxcyoCfwLSNGlF5UB4TkfdFKyKc+6VCnzu7zjUadzJRzlzXXenC7CANUCmnokKmhUV+avAQ7xukty",
	"+arM392hkPBIlXJZaiGtFH3iGdfkTiaysNCxPtks3d7cLCipFjddmTApt6zHo0joWyGR7G6/Dq1pC1hx",
	"WfcSFq5Z0KgmRpUHV+KIQcthFoyvTUzFqRu4livi51KMo6oTca70R0zBafNkaWHPus+Vjn+0rlckpFaI",
	"/606Hz9uzzaZShf1cD/uHX8f8Pd1+Ycl/ZJJN03nTtnUVTZ3TL8tgy6CluzJky0tp5aRzKK91opbFXqL",
	"EyYoGR89EwGzYhehO8v9Emy0NoAzfdR+nSuEmdrvl3+Qt9HRtFoRgs5OoK0/SVG/pPl5eYWy9ROSVPfE",
	"XoFWldbjqFCppFCr/McLmrekf+nMLRuFZUsua2ka4S25U6Ug7lsZCM1hLz/xajoyWti3dpHSwube7Cnn",
	"9eW41kFoVb0T4FUKKRWwXr860zsX6+zSVt+2VLu47aePp2Dj0woS6FGEbOaeM+/8XeouGijr++JK+vZQ",
	"zrfWTLXwSitWj0V+ZNtFc8E1YA3G2WLlWaZgNVaipXd8NSr7vCGtWge16EUtqhdr8fVomWFBAt0QEOTw",
	"W4qory+7Rohr2chmI9YdASWnxye+g+cP+/Ezp4w1ckRvqyiipzatmxL2hAvvS8gew/tmBFlzhaR0hWdS",
	"UthrSEo+b1Axt3yOCqdRqc3a2xKbl4QFc14wJzlhstbtTWmL9W5jGpIzTSdnamegJKUyLQtcktmtNGo0",
	"OUKhbvCSGgo1J2fPHwIEdsW5KiZPe+Lo7NppKIxYlYCxZ3J5i9nRw/yBWgkamu5mZCo4FycvamGghX8k",
	"3WxZOIpWcXeFT1UuT3R9Zucomm4FMfW1wSpYhJP14PPi5Nd6uDn3vQVellRLXRGHL+fFgyt+UoDeoAyF",
	"bDM1zyoNqF2K2oGvu2vqxCfiXhwxBXUvmSrwZBeVMZxVsDAdsEnSrNrAdPRy8QV60Wij57I4eCGng9UD",
	"wKQOfckurqK6zSGP36yb72kbkEvHXTQo+cLe/nxEvkdFI9Kl52yGfOg/DMTViFCrtzPXzxHuNy+4XkZt",
	"z1m0sbbVS1hOlNM03MvdaIf5cgs1kmDlxyX8T3CvPv4nwlfh+WbBrZEFSO87z/e1jQXIDL9ybIXxO6FW",
	"f/9yBs7igrdDssmQbNJnsokipJqwV+5y/l5vRVP2+UqWbESXLJ/GD6Bng1NEmeXyPRBHbydgC2BRf/oG",
	"L29QSCdgaxZ/5fh37Bslxii7/5IRZuS1SjL4JqXLL/tao537HzNh3aIsVPtaf6MvcnRoe3B0CtC9EEyV",
	"QEHzrBld+7/PNUeGGYTvcRk9lIv3JVSi8lQ5yWGN6YjktqFLvYtON/81O+KtPIxsM1k3I/eA7Fjot/Wr",
	"vQv5vQoXIwM1zRwMpX9PBR1v5WpouKIfLeMn9hZ+dDSXwjkfLHIhdrhwLfy/+gHyonKE2FtOfMJmT6DA",
	"6EuAPPA5+RH86q88O84l4KVoRjeUBuHOdKppJofgrLndk0PwPq4lGUExDsESeYjwra4F8V0gwsLA9/zQ",
	"EnEu4Po2ckDgrJZ844YZRE9Uz4w6uxvA+Q2ytiebuc7d3d1NIP+ZdWwavRtOjw73Dj6fHfB3pBmKgtI8",
	"jWM0Ht0iEooBiG+exiM2Whjg0c7o/WSTEwwgveEyUf+K63rX3I7Go6UoNsasG//u0B7tjD4hKm12RmJE",
	"UBj4bGjs+e3NzXjikQgzypla/wpFD4S8NrnYOboeMDe9nw7OwZe/AUb2iVdJc0XVPP6DPHAQCz1ccoGW",
	"r+a9YnANCXQRRSTktwXq+pU+MlUYLG7XW+OacTEpDhKRGpXl+/z7XSWjRsfxTOTk4Ojg/ICxBXz5W4Yz",
	"0W/TTBJeAV9MCUtPgrKejOzFVcG6EpSxMmKK7uk0cCD2/syr04aI/rKiC+sndejZ6sS5Af6GHkBOphgK",
	"M/8B0psUXHOZx6nl4+B9NR4FfqiZzBM/zM3m7ysU0o++/VA0P+kjGIXT0/SFb5mLsTOSsaWp10YQzwlQ",
	"J/Dky9l5PUFeT0Gnxfsmudvr66lDfC9dnxiqXn6/jqIUM6gEZtPL+d++GhkSN/ZWksrczEwcpOJo3FyU",
	"6EeNmTeqP33rTldqU2F5/kC602H/FGXSdk49OdDSOCqyacZISvJn0Fquo8edINzUjSPUDVFOBA17BDpN",
	"vxtDX/7umwbAVzzi7vwGvjHXCf7l2Tkg4htExFhoDYNivGPcAzC20ftuoFJ7Eqf2QaUGSBPFDnlY+X84",
	"maPo7FKH+MNInKLFl0CEGHpBHR1HQTrWsln+pkZeB0h6Pkiq1rYQU1Re2L5EQc4wRV0v2tPC+I3lnpfl",
	"z6zBed39PmKc6mnDakfuTLDajNsWzVtu0K1nuesZbjm5Or/IzAx3qNSJTGgVOr7QoZl/kczRGt5EfCXF",
	"ep6DXvSa6su0uJBtlRLtSW+eJRVozelVVScr1U7uYHwVYXITGqnSxPrDMwe/e7pxttLYCuYNCi0rdNFU",
	"r6njBdNpQvEb64YRXJgytbE+LFvjw55PEC9m3RtKJF1WOaKfmjyEnN8g8GHJ4aJOrLjNsLtFj2970fi7",
	"QJGYuYVoUszmAWJkqegCaNKZ7wFwGqqZQTSarYdGs2dAo9k6aDQzgkaz50SjWZdoNBvQqDUazbpGo1mf",
	"aDTrF40I9KzZcu0l1Cn0ZstnWE5l+98OoXRLrqieR0OUqsGHjrHqVGVIJ5ClMn1ArjbIlZOUTvArKw19",
	"wFgrlWyMZlXXVEZX91WosChN1EuQel++RHRNrYxq9GuC1xGZN6llLWWk+O4LzYUHteSF34nQo9Aody0Z",
	"ER3RYokAxdcgDGJUdpFI/b0bSXC68o+yl7WUTqpRVehVDcxqQKHv8tbVoMMe5TRE27P8JUBNvaqsBK7r",
	"QWXvmDHqKjXQTsMAVegqqTf+hAB78pVGSVdrgkDvrlRkFc37VRWMrLSab9gHewWo0Y32JD+0sfj7qMu4",
	"SD3trymv3SPEc6BDT8BQy5kY0OEZ0aH7Hlb1bV1PR1Iik/5Ooigduj1rglP3wJ7eYVKZJJx0vBH2if8e",
	"2q8mN9g4RzVJwoUu3X6iMQNWDljZlX+HA8v2XYi9Sp1PngRhXPqL37WsXiMY+L7DfKbDk7Ti7/jSgyG4",
	"41WARamiqKH4ehKpRV69pz6wHAb7vKk3CinZ2amFH4fxSwNwvIklWK2jjhm9MHnMsXlvm4Vf659urD3I",
	"jpZTrY41NmbfEJV9G1FZA2cYC0Sw31VKtUabRrxuTixmpudVnlZcn4+1XIjXc0hxcCMywtA2VNutu1CS",
	"p7F+AktPToCpvJUK8z6o2wsLPBo04h2Z7kJBMoAh0xC71hyS6qBgiF32XC2VPcPuHiSvPBJYxqZyM/vt",
	"LHp6UPaXHjlTJ5lCXUUAgqANfM95AJBSgq9XFIWpEaunEBTSvgzZN0GsCyVgDQ8G7vXKPHZDbNlogT2c",
	"FgGuciAP3RDvp+905Ehmu1Z5lHweiyGvGpi+Htf/LFPKWkMyp5eHmbGtq5oZXg0nx2WvTjO363p3ufkz",
	"7OZVyn5zRU82CQozw3EQ7dNUrcLiTZhest+S3Q0zOW/prpYur+1Nb6VctRea6WPysX6cIRaTzgxEMlml",
	"s2hI2PsTdIMyXuiZDXuG7XqkqoG2W/IjbayVLG9r26lUpkxbqDrK1xhwXN/D1OczUAtijtPnOwIZV6bQ",
	"yAFV3ixDmMpBmMMYidS6IJMOb/A2Zf3NzOa6GizPmGEVLhPtNTR3iuwlsqoOrLKHJBWpGf9PuXFgL/s8",
	"zpoS/naQDs+MS6pnndY/1XdjcFaL2Th9lP6o77vqBK17E6OoTv2p71rDnkm7ulKsQqf4j6NdXVbpzaib",
	"vlCv8tBaVjYjpubsrSqKndneVjrfGBzj3aqy6sYuYE9UAUS0k9eL1Y13zczY2ZgHWsv6ljforlqLC/vK",
	"rW81I9noylAmQlw2d0akuzfJNifUhVZt2Hpu2SO3rENuO7sly9e6piqVIcPGqYaeNUcUp3Q5yK8Ur1RO",
	"p6/lHqdkyOo4RYs5TmSwNwmTpo/8nwbmxulwVRY6GocsnbJ1xbgfETYkvcW2xRkWSI175JQujeKf2xgX",
	"x9QqKBIe03alTKXWxoyK+me7wNGVfZoAfjggBHfYcUROVwAJxXMcQIoA9i49eoOEYgO84KcZkQevHQQW",
	"GDk2wCEIEQXUB+dkheqoft9V1YQxM19brXQCii3ecKrpObDDrHqpZ0+amOseKsOVA0M9kXwF2PEMuNEx",
	"ZJS7GQNuPANu9HRuqsahqbb+kOmyeTqt6MJPWgvAOkV79bd6pWNAneJZ2fmSPr7WgyPr8rXoWInGpxvg",
	"cYBHwy7fAjsUkbK1lKRovJgLdBz/DkDPnvoE2Mh7mIDTlYNCAAkC6B7NVxTZ3DMKCPYJpg/AJzYiE3B+",
	"w7wiElJAVg5ibbmQzm+Eb6XWnFlAii49XzhUAZz/huikDrT8yofT45pLEOxotSUmp2SdJagPUPBqVlhi",
	"RpVbiZssraL57mtVtYjJlYveG0KOXlGjM8CoWGUNqPFGHAgFRaru8m67upIUw8zCKhb+rpdUtdCrG0hX",
	"rqCvWj6pc1QTqKSL6N/MwqkmJ8sXTJFDJvFnALo/PNCtr+W1K/ZxKeywVp+2ZzX3sE8QAWK7Ny7Q9xFT",
	"cFpdoa9iUKbdIxNF+XR8Gna5X9Eut6lae5JIde1xlKji+ggUr4msuAQcWwoVp1slS6i4lh57vFbeykn0",
	"ZlQxji3RegyonOiG2VF8RcvSknCLtm8Darya6It2vqeP2q8b5qhplKY3D0CPDA0luAcUeU4E6Qs8KkIv",
	"A4K8zQVKMYRo+1j0eFtfqUjZzPhOBQrVtTdlANb6NxldRn40s/wWw0AGWF4eItK7mEPEaADkOoD8DJgS",
	"nXM3coVTEZKI3JI/LogoPG6DH0MlzAE6Xh50IM8OfOzRSvCIHwTYM+CXHESt/XHxJMP4NogSM3HAlAFT",
	"XhKmLBvthWnQ4VOP22OtRrdW7OwEEUs4AyBAxJLWFeDTx1Ow8WkFCfQoQjb4iOkppOgdA93roG1c7VNv",
	"G3N6nPrUxXZdm3kbom8Dur6M6Nsno5uXJWr3LEE4o/D5DAbMXdeAHb9kA+Z2aMCOmQGTkkcMWK/jZ7Ze",
	"xy/CermD9Rqs10uxXsc9WK/jF2m93Fdgvbq7SrNIHt7CvZrd8L1NWGe4iXPA775iO6tgUYkQqxAREDiQ",
	"H6CCFGA3cJDLpA7QGxyC2/miDm5cBIs3gw0p28oV/CJYDAr8atL/7mEQNFj3/YM93tcq714Qq5fGz0Qp",
	"BAQ5zDFiKny/G9SqvFAyJNOLLU7KuGIyNg1Z/K8oiz+RODNrCSFVXS8bypSxOQy50HGsOXKcksrc7BnA",
	"ngHQtgkKw+okW/bKHnKcfnLzk0EYSsVP2tOn26c8G4qnKuyaPqafG6Sox8LSmT1LJ6x8Ks3Lf4+yb1Ls",
	"i43Ym5b9LnukKoa+X/IzbWyaInVr2zRJskzbtHoK2RiJVsGirFgQY7MNVsEirCq6cREserFcF8HCjMla",
	"BQutrXqja8CrNqIxfVwFi/pmibGuI4O0ChZFE2VEdvuQWyMiW2hnhthF0x7Fsq3tkPixhUmJhWldY8IF",
	"xrAVKdSianCgyA0cSEtOKItKBPFzZf7eefRM1wYjptPSYMRDyZqJuNn11e2qPt+nj/GnenicdNIUINO0",
	"QS0jTAlFHwKxhizo8NecQHSIdor4aCFPeqIh7ilztwbwpfOzLvCVS2tbtauXaRQTMplVVNKfWvp4gogl",
	"qpr//csZOEOUYm8ZNqkbUz0qs5jdKo+nmE2D4pYproEMEXXiTOlvXXmvodHKNn2hF1M3v0LOoejcj5G7",
	"3taZkdvIeTRmUxiuGs7G9FH5s6aDI02AOZhVGFHGIHPS05vkrCs0WgR9NckveRHTY2nmsaaAmp3YdRBV",
	"nby1IbWmbD89Pf1/AAAA//8IuOU2KnsBAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
