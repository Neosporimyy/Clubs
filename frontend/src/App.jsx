import React, { useState, useEffect } from "react";
import ControlPanel from "./components/ControlPanel";
import DataTable from "./components/DataTable";
import FormModal from "./components/FormModal";
import { fetchClubs, fetchClubsByTownName, fetchTowns, createClub, updateClub, deleteClub, createTown, updateTown, deleteTown, fetchClubById, fetchTownById } from "./api";

export default function App() {
  const [search, setSearch] = useState("");
  const [mode, setMode] = useState("clubs");
  const [searchType, setSearchType] = useState("club");
  const [data, setData] = useState([]);
  const [formMode, setFormMode] = useState(null);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const limit = 15;

  const loadData = async (newMode = mode, newSearch = search, newPage = page) => {
    try {
      let result;
      if (newSearch) {
        if (searchType === "club") {
          result = await fetchClubs(newSearch, newPage, limit);
        } else if (searchType === "town") {
          result = await fetchClubsByTownName(newSearch, newPage, limit);
        }
        setMode("clubs");
      } else {
        if (newMode === "clubs") {
          result = await fetchClubs("", newPage, limit);
        } else if (newMode === "towns") {
          result = await fetchTowns("", newPage, limit);
        }
      }

      const processedData = Array.isArray(result)
        ? result
        : result.clubs || result.towns || result.data || result || [];

      const total = result.total || processedData.length;
      const pages = result.pages || Math.ceil(total / limit);

      setData(processedData);
      setTotalPages(pages);
    } catch (error) {
      console.error(error);
      alert('Ошибка загрузки данных');
      setData([]);
    }
  };

  useEffect(() => {
    loadData();
  }, []);

  const handleShowClubs = () => {
    setMode("clubs");
    setSearch("");
    setPage(1);
    loadData("clubs", "", 1);
  };

  const handleShowTowns = () => {
    setMode("towns");
    setSearch("");
    setPage(1);
    loadData("towns", "", 1);
  };

  const handleSearch = () => {
    setPage(1);
    loadData(mode, search, 1);
  };

  const handleSubmitForm = async (formData) => {
    try {
      if (mode === "clubs") {
        if (formMode === "create") await createClub(formData);
        if (formMode === "edit") await updateClub(formData.id, formData);
        if (formMode === "delete") await deleteClub(formData.id);
        if (formMode === "get") {
          const item = await fetchClubById(formData.id);
          setData([item]);
          setTotalPages(1);
          setPage(1);
        }
      } else if (mode === "towns") {
        if (formMode === "create") await createTown(formData);
        if (formMode === "edit") await updateTown(formData.id, formData);
        if (formMode === "delete") await deleteTown(formData.id);
        if (formMode === "get") {
          const item = await fetchTownById(formData.id);
          setData([item]);
          setTotalPages(1);
          setPage(1);
        }
      }
      setFormMode(null);
      if (formMode !== "get") loadData();
    } catch (error) {
      console.error(error);
      alert('Ошибка операции');
    }
  };

  const handlePageChange = (newPage) => {
    setPage(newPage);
    loadData(mode, search, newPage);
  };

  return (
    <div style={{ padding: "30px", fontFamily: "Arial" }}>
      <h2 style={{ textAlign: 'center' }}>Спортивная статистика</h2>

      {formMode && (
        <FormModal
          mode={mode}
          action={formMode}
          onSubmit={handleSubmitForm}
          onCancel={() => setFormMode(null)}
        />
      )}

      <ControlPanel
        search={search}
        setSearch={setSearch}
        searchType={searchType}
        setSearchType={setSearchType}
        onSearch={handleSearch}
        onShowClubs={handleShowClubs}
        onShowTowns={handleShowTowns}
        onCreateClub={() => {
          setMode("clubs");
          setFormMode("create");
        }}
        onEditClub={() => {
          setMode("clubs");
          setFormMode("edit");
        }}
        onDeleteClub={() => {
          setMode("clubs");
          setFormMode("delete");
        }}
        onGetClub={() => {
          setMode("clubs");
          setFormMode("get");
        }}
        onCreateTown={() => {
          setMode("towns");
          setFormMode("create");
        }}
        onEditTown={() => {
          setMode("towns");
          setFormMode("edit");
        }}
        onDeleteTown={() => {
          setMode("towns");
          setFormMode("delete");
        }}
        onGetTown={() => {
          setMode("towns");
          setFormMode("get");
        }}
      />

      <DataTable mode={mode} data={data} />

      <div style={{ display: 'flex', justifyContent: 'center', marginTop: '10px' }}>
        <button onClick={() => handlePageChange(page - 1)} disabled={page <= 1}>↑ Пред</button>
        <span style={{ margin: '0 10px' }}>Страница {page} из {totalPages}</span>
        <button onClick={() => handlePageChange(page + 1)} disabled={page >= totalPages}>↓ След</button>
      </div>
    </div>
  );
}