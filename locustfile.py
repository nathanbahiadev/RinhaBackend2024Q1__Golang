import random
import string

from locust import HttpUser, task


def gerar_transacao():
    return {
        "valor": random.randint(1, 100),
        "tipo": "c",
        "descricao": "abcdefghi"
    }

def selecionar_cliente():
    return random.randint(1, 5)


class RinhaBackendSimulador(HttpUser):
    @task
    def cadastrar_transacao(self):
        self.client.post(f"/clientes/{selecionar_cliente()}/transacoes", json=gerar_transacao())

    @task
    def consultar_extrato(self):
        self.client.get(f"/clientes/{selecionar_cliente()}/extrato")
