import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Turmas() {
  const [turmas, setTurmas] = useState([]);
  const [novaTurma, setNovaTurma] = useState({ nome: '', semestre: '', ano: '', professorID: '' });

  useEffect(() => {
    axios.get('http://localhost:8080/turmas')
      .then(response => setTurmas(response.data))
      .catch(error => console.log(error));
  }, []);

  const handleChange = (e) => {
    setNovaTurma({ ...novaTurma, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/turmas', novaTurma)
      .then(response => {
        setTurmas([...turmas, response.data]);
        setNovaTurma({ nome: '', semestre: '', ano: '', professorID: '' });
      })
      .catch(error => console.log(error));
  };

  return (
    <div>
      <h2>Turmas</h2>
      <ul>
        {turmas.map(turma => (
          <li key={turma.ID}>{turma.Nome} - {turma.Semestre} - {turma.Ano}</li>
        ))}
      </ul>

      <h3>Cadastrar Nova Turma</h3>
      <form onSubmit={handleSubmit}>
        <input type="text" name="nome" placeholder="Nome" value={novaTurma.nome} onChange={handleChange} required />
        <input type="text" name="semestre" placeholder="Semestre" value={novaTurma.semestre} onChange={handleChange} required />
        <input type="text" name="ano" placeholder="Ano" value={novaTurma.ano} onChange={handleChange} required />
        <input type="text" name="professorID" placeholder="ID do Professor" value={novaTurma.professorID} onChange={handleChange} required />
        <button type="submit">Cadastrar</button>
      </form>
    </div>
  );
}

export default Turmas;