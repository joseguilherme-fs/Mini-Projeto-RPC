### 2. Discussões sobre Limitações

* **Escalabilidade:** A escalabilidade do sistema é **baixa**. Ele opera em um único processo de servidor, o que o limita aos recursos de uma única máquina. O uso de um mutex global para todas as operações (Append, Remove, etc.) significa que apenas uma requisição de cliente pode ser processada por vez. Com muitos clientes simultâneos, eles formarão uma fila aguardando a liberação do lock, o que é ruim para a perfomance.

* **Disponibilidade:** A disponibilidade também é **baixa**. O servidor é um ponto único de falha. Se o processo do servidor travar ou a máquina em que ele está rodando falhar, o serviço fica completamente indisponível até que seja reiniciado manualmente.

* **Consistência:** A consistência é **forte**. Graças ao mutex global, todos os clientes, ao acessarem o servidor, terão a mesma visão do estado das listas, garantindo que as leituras sempre reflitam a última escrita concluída.

---

### Pontos de Falha e Melhorias

* **Pontos de Falha e Tratamento:** O principal ponto de falha é o **processo do servidor**. O sistema lida com essa falha através da persistência em disco. Ao reiniciar, ele recupera seu estado carregando o último *snapshot* e reaplicando as operações do *log*. Isso garante a durabilidade dos dados, mas não garante a disponibilidade (o serviço fica fora do ar durante a falha).

* **Como Melhorar a Escalabilidade:** A escalabilidade poderia ser melhorada de duas formas principais:
    1.  **Replicação:** Em vez de um único servidor, teríamos múltiplos servidores (réplicas), cada um com uma cópia dos dados. 
    2.  **Particionamento:** Dividir os dados entre múltiplos servidores. Por exemplo, listas de A-M ficariam no Servidor 1 e listas de N-Z no Servidor 2.

* **Impacto das Melhorias:** A introdução de replicação ou particionamento teria um impacto direto na consistência e disponibilidade:
    * **Disponibilidade aumentaria drasticamente**, pois se uma réplica falhar, as outras podem assumir o tráfego.
    * **Consistência forte se tornaria muito mais difícil de garantir**. Manter todas as réplicas perfeitamente sincronizadas em tempo real introduz latência e complexidade.
