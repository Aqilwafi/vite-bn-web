import React from "react";
import { Create, SimpleForm, TextInput, BooleanInput, NumberInput, DateTimeInput } from "react-admin";
import MDEditor from "@uiw/react-md-editor";

const ArtikelCreate = (props) => {
  const [konten, setKonten] = React.useState("");

  return (
    <Create {...props}>
      <SimpleForm
        // sebelum submit, masukkan konten_md dari state MDEditor ke form values
        onSubmit={(values) => ({ ...values, konten_md: konten })}
      >
        <TextInput source="judul" />
        <TextInput source="slug" />
        <div style={{ marginTop: "1em", marginBottom: "1em" }}>
          <label style={{ display: "block", marginBottom: 4 }}>Konten Markdown</label>
          <MDEditor value={konten} onChange={setKonten} />
        </div>
        <TextInput source="ringkasan" multiline />
        <TextInput source="gambar" />
        <TextInput source="kategori" />
        <TextInput source="penulis" />
        <NumberInput source="waktu_baca" />
        <NumberInput source="jumlah_komentar" />
        <BooleanInput source="unggulan" />
        <DateTimeInput source="tanggal_dibuat" />
        <DateTimeInput source="tanggal_diperbarui" />
      </SimpleForm>
    </Create>
  );
};

export default ArtikelCreate;
