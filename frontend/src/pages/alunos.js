import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Alunos() {
  const [alunos, setAlunos] = useState([]);
  const [novoAluno, setNovoAluno] = useState({ nome: '', matricula: '', turmas: [] });
  const [turmas, setTurmas] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/alunos')
      .then(response => setAlunos(response.data))
      .catch(error => console.log(error));

    axios.get('http://localhost:8080/turmas')
      .then(response => setTurmas(response.data))
      .catch(error => console.log(error));
  }, []);

  const handleChange = (e) => {
    setNovoAluno({ ...novoAluno, [e.target.name]: e.target.value });
  };

  const handleTurmasChange = (e) => {
    const selectedTurmas = Array.from(e.target.selectedOptions, option => parseInt(option.value, 10));
    setNovoAluno({ ...novoAluno, turmas: selectedTurmas });
  };  

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/alunos', novoAluno)
      .then(response => {
        setAlunos([...alunos, response.data]);
        setNovoAluno({ nome: '', matricula: '', turmas: [] });
      })
      .catch(error => console.log(error));
  };

  return (
    <div className="container">
      <h2 className="my-4">Alunos</h2>

      <h3 className="mb-3">Cadastrar Novo Aluno</h3>
      <form onSubmit={handleSubmit} className="mb-4" noValidate>
        <div className="mb-3">
          <label htmlFor="nome" className="form-label">Nome do Aluno</label>
          <input
            type="text"
            className="form-control"
            name="nome"
            id="nome"
            placeholder="Nome do Aluno"
            value={novoAluno.nome}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="matricula" className="form-label">Matrícula</label>
          <input
            type="text"
            className="form-control"
            name="matricula"
            id="matricula"
            placeholder="Matrícula"
            value={novoAluno.matricula}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="turmas" className="form-label">Turmas</label>
          <select
            multiple
            className="form-control"
            name="turmas"
            id="turmas"
            value={novoAluno.turmas}
            onChange={handleTurmasChange}
            required
          >
            {turmas.map(turma => (
              <option key={turma.ID} value={turma.ID}>
                {turma.Nome} - {turma.Semestre}/{turma.Ano}
              </option>
            ))}
          </select>
        </div>
        <button type="submit" className="btn btn-primary">Cadastrar</button>
      </form>

      <h3 className="mb-3">Lista de Alunos</h3>
      <ul className="list-group">
        {alunos.map(aluno => (
          <li key={aluno.ID} className="list-group-item">
            {aluno.Nome} - Matrícula: {aluno.Matricula}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Alunos;