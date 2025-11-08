import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

function ArtikelDetail({ artikel }) {
  return (
    <div className="prose mx-auto px-4 py-10 bg-white rounded-lg shadow-sm">
      <h1 className="text-3xl font-bold mb-6">{artikel.judul}</h1>

      {artikel.gambar && (
        <img
          src={artikel.gambar}
          alt={artikel.judul}
          className="w-full rounded-lg mb-6"
        />
      )}

      <div className="text-sm text-gray-500 mb-4">
        <span>✍️ {artikel.penulis || "Admin"}</span> •{" "}
        <span>
          {new Date(artikel.tanggal_dibuat).toLocaleDateString("id-ID", {
            day: "numeric",
            month: "long",
            year: "numeric",
          })}
        </span>
      </div>

      {/* ✅ Bungkus ReactMarkdown di dalam div agar tidak error */}
      <div className="prose max-w-none">
        <ReactMarkdown remarkPlugins={[remarkGfm]}>
           {artikel.konten_md.replace(/\\n/g, "\n")}
        </ReactMarkdown>
      </div>
    </div>
  );
}

export default ArtikelDetail;
