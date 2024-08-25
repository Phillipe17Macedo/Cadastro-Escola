import React, { useState, useEffect } from "react";
import axios from "axios";

function Atividades() {
  const [atividades, setAtividades] = useState([]);
  const [novaAtividade, setNovaAtividade] = useState({
    nome: "",
    valor: "",
    data: "",
    turmaID: "",
  });
  const [turmas, setTurmas] = useState([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/atividades")
      .then((response) => setAtividades(response.data))
      .catch((error) => console.log(error));

    axios
      .get("http://localhost:8080/turmas")
      .then((response) => setTurmas(response.data))
      .catch((error) => console.log(error));
  }, []);

  const handleChange = (e) => {
    setNovaAtividade({ ...novaAtividade, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const atividadeData = {
      nome: novaAtividade.nome,
      valor: parseFloat(novaAtividade.valor), // Converte o valor para float
      data: novaAtividade.data, // Data no formato "YYYY-MM-DD"
      turmaID: parseInt(novaAtividade.turmaID, 10), // Converte o ID da turma para número
    };

    axios
      .post("http://localhost:8080/atividades", atividadeData)
      .then((response) => {
        setAtividades([...atividades, response.data]);
        setNovaAtividade({ nome: "", valor: "", data: "", turmaID: "" });
      })
      .catch((error) => console.log(error));
  };

  return (
    <div className="container">
      <h2 className="my-4">Atividades</h2>

      <h3 className="mb-3">Cadastrar Nova Atividade</h3>
      <form onSubmit={handleSubmit} className="mb-4" noValidate>
        <div className="mb-3">
          <label htmlFor="nome" className="form-label">
            Nome da Atividade
          </label>
          <input
            type="text"
            className="form-control"
            name="nome"
            id="nome"
            placeholder="Nome da Atividade"
            value={novaAtividade.nome}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="valor" className="form-label">
            Valor
          </label>
          <input
            type="number"
            step="0.01"
            className="form-control"
            name="valor"
            id="valor"
            placeholder="Valor da Atividade"
            value={novaAtividade.valor}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="data" className="form-label">
            Data
          </label>
          <input
            type="date"
            className="form-control"
            name="data"
            id="data"
            value={novaAtividade.data}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="turmaID" className="form-label">
            Turma
          </label>
          <select
            className="form-control"
            name="turmaID"
            id="turmaID"
            value={novaAtividade.turmaID}
            onChange={handleChange}
            required
          >
            <option value="" disabled>
              Selecione uma Turma
            </option>
            {turmas.map((turma) => (
              <option key={turma.ID} value={turma.ID}>
                {turma.Nome} - {turma.Semestre}/{turma.Ano}
              </option>
            ))}
          </select>
        </div>
        <button type="submit" className="btn btn-primary">
          Cadastrar
        </button>
      </form>

      <h3 className="mb-3">Lista de Atividades</h3>
      <ul className="list-group">
        {atividades.map((atividade) => (
          <li key={atividade.ID} className="list-group-item">
            {atividade.Nome} - {atividade.Valor} pontos - {atividade.Data} -
            Turma:{" "}
            {atividade.Turma
              ? `${atividade.Turma.Nome} - ${atividade.Turma.Semestre}/${atividade.Turma.Ano}`
              : "Sem Turma"}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Atividades;
