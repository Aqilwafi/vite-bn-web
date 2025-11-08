import { Search } from "lucide-react";

export default function SearchBar({ value, onChange }) {
  return (
    <div className="flex items-center border rounded-full px-3 py-1.5 bg-gray-50 focus-within:ring-2 focus-within:ring-green-500">
      <Search className="w-5 h-5 text-gray-500" />
      <input
        type="text"
        placeholder="Cari artikel..."
        value={value}
        onChange={(e) => onChange(e.target.value)}
        className="ml-2 bg-transparent outline-none text-sm text-gray-700 w-28 focus:w-44 transition-all duration-300"
      />
    </div>
  );
}
