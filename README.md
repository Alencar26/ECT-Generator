# 🧾 ECT — Gerador de Evidências de Teste

O **ECT (Evidências de Casos de Teste)** é uma ferramenta simples e leve para **gerar relatórios em PDF** a partir de **imagens ou capturas de tela** (screenshots) contidas em um diretório.

Seu objetivo é **padronizar a geração de evidências visuais** dos testes de software manuais, garantindo clareza, rastreabilidade e uma apresentação profissional das execuções.

---

## ⚙️ Funcionamento

O software percorre o diretório atual, **coleta todas as imagens** (`.png`, `.jpg`, `.jpeg`, etc.) e **gera automaticamente um PDF** com capa e todas as imagens ordenadas na sequência encontrada.

A capa contém informações básicas como:
- Nome do autor (extraído do nome do binário)
- Identificador do relatório (também no nome do binário)
- Data e hora da geração

---

## 🚀 Como usar

1. Coloque o binário do **ECT** no diretório onde estão suas imagens.
2. Garanta que as imagens estejam com o nome na ordem desejada (ex: `01_login.png`, `02_cadastro.png`...).
3. Execute o binário:

   ```bash
   ./ECT_0001_Andre_Alencar
   ```

4. O programa irá gerar um arquivo chamado:

   ```
   ECT_0001_Andre_Alencar.pdf
   ```

   no mesmo diretório.

---

## 🧱 Estrutura esperada

```
/meu_teste/
├── 01_login.png
├── 02_cadastro.png
├── 03_relatorio.png
└── ECT_0001_Andre_Alencar
```

Ao executar o binário, será criado:

```
ECT_0001_Andre_Alencar.pdf
```

---

## 🧩 Convenção do nome do binário

O nome do binário deve **seguir o padrão**:

```
ECT_<número>_<NomeDoAutor>
```

Exemplos válidos:
```
ECT_0001_Pedro
ECT_0045_Maria
ECT_0123_Felipe
```

Essa convenção é usada para gerar automaticamente o nome e os dados da capa do PDF.

---

## 📦 Distribuição

O software é distribuído como **um único executável** (sem dependências).
Basta copiar o binário para o diretório desejado e executá-lo.

Releases oficiais estão disponíveis na aba [**Releases**](../../releases) deste repositório.

---

## 🧰 Requisitos

- Sistema operacional: **Linux**, **Windows** ou **macOS**
- Nenhuma instalação adicional é necessária

---

## 🧑‍💻 Exemplo de uso prático

Imagine que você finalizou os testes de um módulo de login e salvou as evidências:

```
login_test/
├── 01_acesso.png
├── 02_preenchimento.png
├── 03_resultado.png
└── ECT_0007_JoaoPedro
```

Ao rodar o comando:

```bash
./ECT_0007_JoaoPedro
```

Será gerado o arquivo:

```
Evidencias_0007_JoaoPedro.pdf
```

contendo a capa e todas as imagens organizadas.

---

## 📄 Licença

[MIT](./LICENSE)
---

## ✨ Autores e Colaboradores

- André Alencar

---

> 💡 **ECT**: simples, direto e padronizado — a forma mais rápida de documentar visualmente seus testes manuais.
