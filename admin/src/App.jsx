import { Admin, Resource, ListGuesser, EditGuesser } from "react-admin";
import dataProvider from "./providers/data";
import ArtikelCreate from "./pages/TambahArtikel";
import Dashboard from "./pages/Dashboard";

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
