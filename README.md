## Gemini AI Sandbox

Gemini AI Sandbox is an application designed to manage pairs of questions and answers utilized by Gemini AI to provide responses based on predefined information.

### Tech Used

* [GO](https://go.dev/) as the programming language
* [Gin](https://gin-gonic.com/) as the HTTP Framework
* [MySQL](https://www.mysql.com/) as the Database
* [Generative AI SDK](https://github.com/google/generative-ai-go) package from Google for interacting with Gemini AI

### How to Run

#### Prerequisite

Ensure you have a [Google Cloud Project API Key](https://makersuite.google.com/app/apikey).

#### Running on Docker

> **Note**: Update the `API_KEY` in the [docker-compose.yml](https://github.com/alqinsidev/go-gemini-sandbox/blob/main/docker-compose.yml) with your API key.

Follow these steps to run Gemini AI Sandbox as a Docker container:

1. Start the Docker container stack by executing this command:

    ```bash
    docker-compose up -d
    ```

2. After the container has started, run the database migration with this command:

    ```bash
    docker exec gemini-api "./migration"
    ```

### Populating Information

Before using the [Chat API](#chat), you need to set pairs of questions and answers using the [Informations API](#informations).

To do this, provide a list of questions along with their expected answers.

#### Use Case

For example, if you want Gemini to answer questions about personal information, you need to set up relevant information.

##### Information Example

Store this pair of questions and expected answers using [Informations API](#informations):

**Question**: What is your name?

**Answer**: John Doe

The expected behavior for Gemini when using the [Chat API](#chat) after storing this information would be:

```plaintext
// Question
What is your last name?

// Answer
Doe
```


## API

Here are list of available API

### Informations

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
