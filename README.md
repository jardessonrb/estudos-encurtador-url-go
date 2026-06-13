# estudos-encurtador-url-go

A ideia do projeto é ESTUDAR mais sobre a linguagem Go, mas principalmente estudar conceitos de arquitetura de software usando ferramentas como Kubernetes, Docker, Redis, PostgreSQL.

O projeto é um encurtador de url, com funcionalidades simples, basicamente duas funcionalidades:

- POST: /gerar-codigo
    - Body
    ```
    {
        "url": "https://www.seulinksuper.com/"
    }
    ```

    - Response
    ```
    {
        "id": 13,
        "codigo": "xQWt"
    }
    ```
- GEST: /{codigo}
    - Chamada
    ```
        https://encurtador-url-go/Xdsf
    ```
    - retorno
    ```
    {
        "url": "https://www.seulinksuper.com/"
    }
    ```

Isso é simples, mas a ideia é focar na arquitetura de um solução mais potente no mundo real.

A arquitetura passou a seguir o canal [Renato Augusto no YouTube](https://www.youtube.com/watch?v=m_anIoKW7Jg)

![Arquitetura](/docs/arquitetura.png)
