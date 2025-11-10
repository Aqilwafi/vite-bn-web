import simpleRestProvider from "ra-data-simple-rest";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:4000/api";

const simpleProvider = simpleRestProvider(API_URL);

const dataProvider = {
  ...simpleProvider,

  getOne: async (resource, params) => {
    if (resource === "artikel") {
      const res = await fetch(`${API_URL}/${resource}/id/${params.id}`);
      if (!res.ok) throw new Error("Artikel not found");
      const data = await res.json();
      return { data };
    }
    return simpleProvider.getOne(resource, params);
  },

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
};

export default dataProvider;
