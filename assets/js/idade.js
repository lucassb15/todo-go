
function calcularIdade() {
    let idade = document.getElementById("inputDate").value
    let conteudo = document.getElementById('conteudo')

    let dataAtual = new Date().getFullYear();
    let dataUsuario = new Date(idade).getFullYear();
    let idadeFinal = dataAtual - dataUsuario;
    conteudo.innerHTML = idadeFinal;
}
