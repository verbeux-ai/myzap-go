# Z-API Go

Essa biblioteca foi desenvolvida para o uso do [https://github.com/verbeux-ai/myzap](myzap), api do whatsapp usando a linguagem go.

Status: Em desenvolvimento

## Enviando Mensagem

Este é um exemplo de envio de mensagem a partir do client.

```go
import myzap "github.com/verbeux-ai/myzap"
import "context"

client = myzap.NewClient(
    myzap.WithToken(os.Getenv("TOKEN")),
    myzap.WithSessionKey(os.Getenv("SESSION_KEY")),
    myzap.WithBaseUrl(os.Getenv("BASE_URL")),
)
response, err := client.SendTextMessage(context.Background(), &z_api.TextMessageRequest{
    Number:   "558500000000",
    Phone: "Teste",
})
if err != nil {
    panic(err)
}
fmt.Println(response)
```

## Escutando mensagens

Este é um exemplo de como escutar mensagens no webhook

```go
import "github.com/verbeux-ai/myzap/listener"

whatsappListener := listener.NewMessageListener()
whatsappListener.HandleErrors(func (err error) {
    fmt.Println("fail", err)
})

// register listeners
whatsappListener.OnMessage(func (message *listener.WebhookMessage) error {
  
        // treat your text message here
  
    
    return nil
})

if err := whatsappListener.ReadBodyAsync(ctx.Request().Body); err != nil {
    panic(err)
}
```

## Features disponíveis

| Funcionalidade           | Implementado |
|--------------------------|--------------|
| Enviar Mensagem de Texto | ✔            |
| Cria instancia           | ✔            |
| Deleta instancia         | ✔            |
| Obter Status             | ✔            |
| Obter Imagem do QR Code  | ✔            |
| Obter Tags               | ✔            |
| Criar Tags               | ✔            |
| Atribuir Tags            | ✔            |


> Você está convidado a contribuir ao repositório!