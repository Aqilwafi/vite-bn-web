import { Admin, Resource, ListGuesser, EditGuesser } from "react-admin";
import simpleRestProvider from "ra-data-simple-rest";
import ArtikelCreate from "./pages/TambahArtikel";
import Dashboard from "./pages/Dashboard";

const dataProvider = simpleRestProvider("http://localhost:4000");

export default function AdminPanel() {
  return (
    <Admin dataProvider={dataProvider} dashboard={Dashboard}>
      <Resource
        name="api/artikel"
        list={ListGuesser}
        edit={EditGuesser}
        create={ArtikelCreate}
      />
    </Admin>
  );
}
