import React from "react";
import { Create, SimpleForm, TextInput } from "react-admin";
import MDEditor from "@uiw/react-md-editor";

const ArtikelCreate = (props) => {
  const [value, setValue] = React.useState("");

  return (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="judul" />
        <TextInput source="slug" />
        <div>
          <label>Konten Markdown</label>
          <MDEditor value={value} onChange={setValue} />
        </div>
      </SimpleForm>
    </Create>
  );
};

export default ArtikelCreate;
