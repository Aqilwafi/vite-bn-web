import simpleRestProvider from "ra-data-simple-rest";

const API_URL = import.meta.env.VITE_API_URL || "https://bnadmin-production-f31e.up.railway.app/api";

const simpleProvider = simpleRestProvider(API_URL);

const dataProvider = {
  ...simpleProvider,

  // GET /api/artikel/id/:id untuk Edit
  getOne: async (resource, params) => {
    if (resource === "artikel") {
      const res = await fetch(`${API_URL}/${resource}/id/${params.id}`);
      if (!res.ok) throw new Error("Artikel not found");
      const data = await res.json();
      return { data };
    }
    return simpleProvider.getOne(resource, params);
  },

  // PUT /api/artikel/id/:id untuk Update
  update: async (resource, params) => {
    if (resource === "artikel") {
      const res = await fetch(`${API_URL}/${resource}/id/${params.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(params.data),
      });
      if (!res.ok) throw new Error("Update failed");
      const data = await res.json();
      return { data };
    }
    return simpleProvider.update(resource, params);
  },

  // ✅ POST /api/artikel/create untuk Create
  create: async (resource, params) => {
    if (resource === "artikel") {
      const res = await fetch(`${API_URL}/${resource}/create`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(params.data),
      });
      if (!res.ok) throw new Error("Create failed");
      const data = await res.json();
      return { data };
    }
    return simpleProvider.create(resource, params);
  },
  
  delete: async (resource, params) => {
  if (resource === "artikel") {
    const res = await fetch(`${API_URL}/${resource}/id/${params.id}`, {
      method: "DELETE",
    });
    if (!res.ok) throw new Error("Delete failed");
    return { data: params.previousData };
  }
  return simpleProvider.delete(resource, params);
},

};

export default dataProvider;
