# MyStore

> The new version of [eStore](https://github.com/tanhaok/eStore)

## Overview and component

|Component| Stand for| Language | Measure | Code Coverage |
| -- | -- | -- | -- | -- |
| `api`| api gateway| Java | | |
| `cms`| content management system -  For managing product information, categories, etc.| Java | | |
| `search` | product search | Java | | |
| `iam`| identify and access management - JWT, OAuth for secure user sessions| Go |[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=anyshop_iam&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=anyshop_iam) |[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=anyshop_iam&metric=coverage)](https://sonarcloud.io/summary/new_code?id=anyshop_iam) |
| `ims`| inventory management system| Go | | |
| `oms`| order management system - Handle order processing, inventory management| Go | | |
| `pgi`| payment gateway integration - PayPal, Stripe, or other payment processors | Rust | | |
| `dms`| delivery management system | Rust | | |
| `map`| marketing and promotion | Rust | | |
| `aar`| analytics and reporting | Python | | |
| `notify`| notifications | Python | [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=anyshop_notify&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=anyshop_notify) | [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=anyshop_notify&metric=coverage)](https://sonarcloud.io/summary/new_code?id=anyshop_notify) |
| `recommend` | recommend system | python | | |
| `support` | customer support | NodeJS | | |
| `feedback`| user feedback | NodeJS| | |
| `shop`| user page| Typescript | | |
| `admin`| admin page| Typescript | | |
