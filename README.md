## Gemini AI Sandbox
Application that managing pairs of Question and Answer that used by Gemini AI to provide answer based on information that we set before.

## Tech Used
* [GO](https://go.dev/) as a programming language
* [Gin](https://gin-gonic.com/) as HTTP Framework
* [MySQL](https://www.mysql.com/) as Database
* [Generative AI SDK](https://github.com/google/generative-ai-go) package from google to interact with Gemini AI

## How to Run

### Prequesite
* [Google Cloud Project API Key](https://makersuite.google.com/app/apikey)

### Running in Docker

> **NOTE**: Make sure you change the `API_KEY` on the [docker-compose.yml](https://github.com/alqinsidev/go-gemini-sandbox/blob/main/docker-compose.yml) into your `API_KEY`

You can run this as a docker container by following this step:

1. Start docker container stack by executing this command

```
docker-compose up -d
```

2. After the container started, run database migration by running this command

```
docker exec gemini-api "./migration"
```

## Populating Information
Before we can use the [Chat API](#chat), you need to set the pairs of question and answer by using [Informations API](#informations). 

The information you need to provide is a list of question and it's expected answer to the question.

### Usecase
You want gemini to answer you personal information, then you need to set a few information about that.

#### Information Example
Store this pair of question and it's expected answer:

**Question**: What is your name?

**Answer**: John Doe

This is expected behavior for Gemini by using [Chat API](#chat) after you store that information

```
// Question
what is your last name?

// Answer
Doe
```

## API

This is list of available API

### Informations

<details>
<summary> 
    Get Stored Information List
</summary>

##### Description

Return stored information

##### URL

`GET /informations`

##### Response - 200

```json
{
    "data": [
        {
            "id": int,
            "question": string,
            "answer": string
        }
    ],
    "success": boolean
}
```

</details>

<details>
<summary> 
    Get Stored Information By ID
</summary>

##### Description

Return stored information by id

##### URL

`GET /informations/:id`

#### Params
`id`: integer

##### Response - 200

```json
{
    "data":{
        "id": int,
        "question": string,
        "answer": string
    },
    "success": boolean
}
```

</details>

<details>
<summary> 
    Add New Information
</summary>

##### Description

Store new pair of question and answer used later by gemini to provide an answer

##### URL

`POST /informations`

##### Request Body
```json
{
    "question": string, // user question
    "answer": string // answer expectation
}
```

##### Response - 200

```json
{
    "data":{
        "id": int,
        "question": string,
        "answer": string
    },
    "success": boolean
}
```

</details>
<details>
<summary> 
    Edit Information
</summary>

##### Description

Edit stored pair of question and answer

##### URL

`PUT /informations/:id`

##### Request Params
`id`: integer

##### Request Body
```json
{
    "question": string, // user question
    "answer": string // answer expectation
}
```

##### Response - 200

```json
{
    "data":{
        "id": int,
        "question": string,
        "answer": string
    },
    "success": boolean
}
```

</details>

<details>
<summary> 
    Delete Stored Information
</summary>

##### Description

Delete stored information

##### URL

`DELETE /informations/:id`

#### Params
`id`: integer

##### Response - 200

```json
{
    "data":{
        "id": int,
        "question": string,
        "answer": string
    },
    "success": boolean
}
```

</details>

### Chat

<details>
<summary> 
    Ask a question
</summary>

##### Description

Ask Gemini a question, and gemini will answer it based on the information stored before

##### URL

`POST /chat`

##### Request Body

```json
{
    "question": string
}
```

##### Response - 200

```json
{
    "data": {
        "question": string,
        "answer": string
    },
    "success": boolean
}
```

</details>
