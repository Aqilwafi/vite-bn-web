import { Admin, Resource, ListGuesser, EditGuesser } from "react-admin";
import simpleRestProvider from "ra-data-simple-rest";
import ArtikelCreate from "./pages/TambahArtikel";
import Dashboard from "./pages/Dashboard";

const dataProvider = simpleRestProvider("http://localhost:4000/api");

function App() {
  return (
    <Admin dataProvider={dataProvider} dashboard={Dashboard}>
      <Resource
        name="artikel"
        list={ListGuesser}
        edit={EditGuesser}
        create={ArtikelCreate}
      />
      
    </Admin>
  );
}

export default App;
