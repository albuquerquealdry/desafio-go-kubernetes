Que tal um desafio que envolva a criação de um aplicativo simples em Go que interaja com o Kubernetes? Aqui está o desafio:

### Desafio: Aplicativo Go para Interagir com Kubernetes

#### Objetivo:
Criar um aplicativo em Go que se conecte a um cluster Kubernetes e execute algumas operações básicas, como listar pods, criar um novo deployment e escalar o número de réplicas de um deployment existente.

#### Requisitos:
1. Conhecimento básico de Go.
2. Configuração de acesso ao cluster Kubernetes (pode ser um cluster local ou em nuvem).
3. Acesso ao pacote cliente oficial do Kubernetes para Go (kubernetes/client-go).

#### Tarefas:
1. **Listar Pods:**
   - Conectar-se ao cluster Kubernetes.
   - Listar todos os pods em um namespace específico.
   - Exibir informações básicas de cada pod, como nome, namespace e status.

2. **Criar Deployment:**
   - Criar um novo deployment no cluster.
   - O deployment deve consistir em pelo menos uma réplica de um contêiner simples (por exemplo, nginx).
   - Especificar o nome do deployment, nome do contêiner, imagem e número de réplicas desejado.

3. **Escalonar Deployment:**
   - Aumentar o número de réplicas de um deployment existente.
   - Permitir que o usuário especifique o nome do deployment e o número desejado de réplicas.

#### Observações:
- Utilize o pacote cliente oficial do Kubernetes para Go (client-go) para interagir com o cluster Kubernetes.
- Certifique-se de gerenciar erros adequadamente em todas as operações.
- Documente seu código de forma clara e concisa.
- Teste seu aplicativo em um ambiente Kubernetes real ou simulado.

Este desafio proporcionará a oportunidade de aprender sobre a integração do Go com o Kubernetes e praticar o desenvolvimento de aplicativos que interagem com um cluster de contêineres. Boa sorte!
