package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-11-20 17:10:43.209Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": true
				}
			},
			{
				"id": "rtnsxvculz748mt",
				"created": "2023-12-02 15:15:41.632Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "settings",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "r79wlvpj",
						"name": "defaultCurrency",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "e1ijbpaf",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ofeakjwc",
						"name": "spendingLimit",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 0,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "byln6tut",
						"name": "avatar",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = user",
				"viewRule": "@request.auth.id = user",
				"createRule": "@request.auth.id != '' && @request.auth.id = @request.data.user",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			},
			{
				"id": "uo6axjjrwj6yudy",
				"created": "2023-12-02 15:28:58.736Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "transactions",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "opcigsck",
						"name": "date",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "gvmsugon",
						"name": "note",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "06ufmhno",
						"name": "amount",
						"type": "number",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": 0,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "ik237gqm",
						"name": "currency",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zzptfxpa",
						"name": "isExpense",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "i3j4scqm",
						"name": "isTransfert",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "sctzyhit",
						"name": "incomeCategory",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "asj1m8bb30303x5",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "asfzrpvu",
						"name": "expenseCategory",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7nxkyapivaoxy6h",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "3wswuycd",
						"name": "description",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": true
						}
					},
					{
						"system": false,
						"id": "7rb6hsmn",
						"name": "picture",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/png",
								"image/jpeg"
							],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "eldbtijy",
						"name": "company",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "z1ts4ey0o1vjc0j",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "nuhmhh9e",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ei4f3hsr",
						"name": "fromAccount",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "r1zqilx88trto77",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ssxricqu",
						"name": "toAccount",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "r1zqilx88trto77",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "wp5s7u2m",
						"name": "instalment",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "2b1nhmjfmpx6okg",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "fvjcicab",
						"name": "fromSaving",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "peuknrhwz9dib6x",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "bmdyuzad",
						"name": "toSaving",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "peuknrhwz9dib6x",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "5tsewt1r",
						"name": "fees",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "nlzyk2g5",
						"name": "repeatable",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "ojtvcr6c",
						"name": "pace",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "mnajtx5s",
						"name": "paceUnit",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"day",
								"week",
								"month",
								"year"
							]
						}
					},
					{
						"system": false,
						"id": "b0dsepcy",
						"name": "next",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "uo6axjjrwj6yudy",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = user",
				"viewRule": "@request.auth.id = user",
				"createRule": "(@request.auth.id != '' && @request.auth.id = @request.data.user) \n&& \n(\n  (\n    (@request.data.isExpense = true) && (@request.data.expenseCategory != '') && (@request.data.incomeCategory = '')\n  )\n  || \n  (\n    (@request.data.isTransfert = true) && (@request.data.expenseCategory = '') && (@request.data.incomeCategory = '') && (@request.data.toAccount != '')\n  ) \n  ||\n  (\n    (@request.data.incomeCategory != '') &&(@request.data.expenseCategory = '')\n  ) \n)",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			},
			{
				"id": "z1ts4ey0o1vjc0j",
				"created": "2023-12-02 15:30:26.972Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "transactionCompanies",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "kc81nrmx",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "cdbxi23b",
						"name": "website",
						"type": "url",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					},
					{
						"system": false,
						"id": "lytaat9w",
						"name": "logo",
						"type": "url",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "7nxkyapivaoxy6h",
				"created": "2023-12-02 15:40:09.615Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "expenseCategories",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "gwlkhnan",
						"name": "icon",
						"type": "file",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/svg+xml"
							],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "rppcodvb",
						"name": "name_en",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "20xuhzqu",
						"name": "name_fr",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "riepgyhl",
						"name": "name_de",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "nsace8gq",
						"name": "emoji",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "vmbmk8ry",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "wqeap3mm",
						"name": "bgColor",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "nw7sgtx6",
						"name": "user",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "user = '' || @request.auth.id = user",
				"viewRule": "user = '' || @request.auth.id = user",
				"createRule": "@request.auth.id != '' && @request.auth.id = @request.data.user",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			},
			{
				"id": "2b1nhmjfmpx6okg",
				"created": "2023-12-02 15:48:08.652Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "instalments",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "uzhqiuue",
						"name": "total",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "r1zqilx88trto77",
				"created": "2023-12-02 15:50:24.480Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "accounts",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "bpk9dy2m",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "jideofms",
						"name": "group",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "6ctn2l54gjz8up4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "1gudwmrc",
						"name": "currency",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mnl7yeyr",
						"name": "creditLimit",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "l2n3jvj9",
						"name": "description",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "sbkvfxdu",
						"name": "includeInTotal",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "v8qi7ioa",
						"name": "hideInSetting",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "j14srt5e",
						"name": "user",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = user",
				"viewRule": "@request.auth.id = user",
				"createRule": "@request.auth.id != '' && @request.auth.id = @request.data.user",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			},
			{
				"id": "asj1m8bb30303x5",
				"created": "2023-12-02 15:57:41.600Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "incomeCategories",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "xnxu6e02",
						"name": "icon",
						"type": "file",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/svg+xml"
							],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "g3lknthd",
						"name": "name_en",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "c9w4h5hg",
						"name": "name_fr",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "vvg1hi9e",
						"name": "name_de",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "oqhttitf",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "rc7sbqau",
						"name": "emoji",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "4w6n1im2",
						"name": "bgColor",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "yleedczp",
						"name": "user",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "user = '' || @request.auth.id = user",
				"viewRule": "user = '' || @request.auth.id = user",
				"createRule": "@request.auth.id != '' && @request.auth.id = @request.data.user",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			},
			{
				"id": "20vk5ys2v509j6m",
				"created": "2023-12-02 16:10:41.680Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "budgets",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "asjbukjm",
						"name": "amount",
						"type": "number",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "tlw0wdtl",
						"name": "currency",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "yydmxxra",
						"name": "date",
						"type": "date",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "mfpo2hfk",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "7ydgfrtp",
						"name": "incomeCategory",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "asj1m8bb30303x5",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "rtkyfz8l",
						"name": "expenseCategory",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7nxkyapivaoxy6h",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "axfvarwm",
						"name": "next",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "20vk5ys2v509j6m",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "sie0woo3",
						"name": "repeatable",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "khzqxsxj",
						"name": "pace",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "bg3pnn9f",
						"name": "paceUnit",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"day",
								"week",
								"month",
								"year"
							]
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = user",
				"viewRule": "@request.auth.id = user",
				"createRule": "@request.auth.id != '' && @request.auth.id = @request.data.user",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			},
			{
				"id": "peuknrhwz9dib6x",
				"created": "2023-12-02 16:16:48.911Z",
				"updated": "2024-01-27 16:37:29.166Z",
				"name": "savings",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "sucef2zf",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "f5ncdhqq",
						"name": "description",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "f9imi5xk",
						"name": "note",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "7mnz3bsx",
						"name": "pictures",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/png",
								"image/jpeg"
							],
							"thumbs": [],
							"maxSelect": 5,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "golohdtc",
						"name": "goal",
						"type": "number",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "bvgdhvpn",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "l7aecqsn",
						"name": "currency",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "uzlhncoy",
						"name": "icon",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "07poxjeo",
						"name": "collected",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "qqlwpo2d",
						"name": "account",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "r1zqilx88trto77",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = user",
				"viewRule": "@request.auth.id = user",
				"createRule": "(@request.auth.id != '') && \n(@request.auth.id = @request.data.user) &&\n(@request.data.account.group.name_en = 'Savings')",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			},
			{
				"id": "6ctn2l54gjz8up4",
				"created": "2023-12-04 18:16:18.731Z",
				"updated": "2024-01-27 16:37:29.167Z",
				"name": "AccountGroups",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "knktbf19",
						"name": "icon",
						"type": "file",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/svg+xml"
							],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "4lzy70rp",
						"name": "name_en",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "fbs4mplh",
						"name": "name_fr",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "ihlhgner",
						"name": "name_de",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zlake7oc",
						"name": "emoji",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "061v4s2r",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "7bvcdko0",
						"name": "user",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "user = '' || @request.auth.id = user",
				"viewRule": "user = '' || @request.auth.id = user",
				"createRule": "@request.auth.id != '' && @request.auth.id = @request.data.user",
				"updateRule": "@request.auth.id = user",
				"deleteRule": "@request.auth.id = user",
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
