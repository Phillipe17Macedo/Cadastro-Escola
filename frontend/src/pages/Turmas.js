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
  
    const turmaData = {
      nome: novaTurma.nome,
      semestre: novaTurma.semestre,
      ano: parseInt(novaTurma.ano, 10),  // Enviar como número
      professorID: parseInt(novaTurma.professorID, 10)  // Também como número
    };
  
    axios.post('http://localhost:8080/turmas', turmaData)
      .then(response => {
        setTurmas([...turmas, response.data]);
        setNovaTurma({ nome: '', semestre: '', ano: '', professorID: '' });
      })
      .catch(error => console.log(error));
  };    

  return (
    <div className="container">
      <h2 className="my-4">Turmas</h2>

      <h3 className="mb-3">Cadastrar Nova Turma</h3>
      <form onSubmit={handleSubmit} className="mb-4" noValidate>
        <div className="mb-3">
          <label htmlFor="nome" className="form-label">Nome da Turma</label>
          <input
            type="text"
            className="form-control"
            name="nome"
            id="nome"
            placeholder="Nome da Turma"
            value={novaTurma.nome}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="semestre" className="form-label">Semestre</label>
          <input
            type="text"
            className="form-control"
            name="semestre"
            id="semestre"
            placeholder="Semestre"
            value={novaTurma.semestre}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="ano" className="form-label">Ano</label>
          <input
            type="number"
            className="form-control"
            name="ano"
            id="ano"
            placeholder="Ano"
            value={novaTurma.ano}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="professorID" className="form-label">ID do Professor</label>
          <input
            type="number"
            className="form-control"
            name="professorID"
            id="professorID"
            placeholder="ID do Professor"
            value={novaTurma.professorID}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">Cadastrar</button>
      </form>

      <h3 className="mb-3">Lista de Turmas</h3>
      <ul className="list-group">
        {turmas.map(turma => (
          <li key={turma.ID} className="list-group-item">
            {turma.Nome} - {turma.Semestre} - {turma.Ano}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Turmas;