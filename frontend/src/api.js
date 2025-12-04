const API_URL = "http://localhost:8080/api";

export async function fetchClubs(search = "", page = 1, limit = 15) {
  const url = search
    ? `${API_URL}/clubs/search?search=${encodeURIComponent(search)}&page=${page}&limit=${limit}`
    : `${API_URL}/clubs?page=${page}&limit=${limit}`;
  const res = await fetch(url);
  if (!res.ok) throw new Error('Ошибка при получении клубов');
  return res.json();
}

export async function fetchClubsByTownName(search = "", page = 1, limit = 15) {
  const url = search
    ? `${API_URL}/towns/search?search=${encodeURIComponent(search)}&page=${page}&limit=${limit}`
    : `${API_URL}/clubs?page=${page}&limit=${limit}`;
  const res = await fetch(url);
  if (!res.ok) throw new Error('Ошибка при получении клубов по названию города');
  return res.json();
}

export async function fetchTowns(search = "", page = 1, limit = 15) {
  const res = await fetch(`${API_URL}/towns?search=${encodeURIComponent(search)}&page=${page}&limit=${limit}`);
  if (!res.ok) throw new Error('Ошибка при получении городов');
  return res.json();
}

export async function fetchClubById(id) {
  const res = await fetch(`${API_URL}/clubs/${id}`);
  if (!res.ok) throw new Error('Ошибка при получении клуба по ID');
  return res.json();
}

export async function fetchTownById(id) {
  const res = await fetch(`${API_URL}/towns/${id}`);
  if (!res.ok) throw new Error('Ошибка при получении города по ID');
  return res.json();
}

export async function createClub(data) {
  const res = await fetch(`${API_URL}/clubs`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error('Ошибка при создании клуба');
  return res.json();
}

export async function updateClub(id, data) {
  const res = await fetch(`${API_URL}/clubs/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error('Ошибка при обновлении клуба');
  return res.json();
}

export async function deleteClub(id) {
  const res = await fetch(`${API_URL}/clubs/${id}`, { method: "DELETE" });
  if (!res.ok) throw new Error('Ошибка при удалении клуба');
  return res.status === 204;
}

export async function createTown(data) {
  const res = await fetch(`${API_URL}/towns`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error('Ошибка при создании города');
  return res.json();
}

export async function updateTown(id, data) {
  const res = await fetch(`${API_URL}/towns/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error('Ошибка при обновлении города');
  return res.json();
}

export async function deleteTown(id) {
  const res = await fetch(`${API_URL}/towns/${id}`, { method: "DELETE" });
  if (!res.ok) throw new Error('Ошибка при удалении города');
  return res.status === 204;
}