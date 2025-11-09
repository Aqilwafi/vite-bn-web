import SearchBar from "./SearchBar";
import { getCurrentDate } from "../utils/dateUtils";

export default function Header({ searchQuery, setSearchQuery }) {
  return (
    <header className="bg-white shadow-sm sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-4 py-4">
        <div className="flex items-center justify-between">
          {/* Logo */}
          <div className="flex items-center gap-2">
            <div className="w-8 h-8 bg-green-600 rounded-full flex items-center justify-center">
              <span className="text-white text-xs">🌿</span>
            </div>
            <span className="font-semibold text-gray-800">Baitun Naim</span>
          </div>

          {/* Navigation + Search */}
          <nav className="hidden md:flex items-center gap-6">
            <a href="#" className="text-gray-700 hover:text-green-600 font-medium">HOME</a>
            <a href="#" className="text-gray-700 hover:text-green-600">PROFIL</a>
            <a href="#" className="text-gray-700 hover:text-green-600">JENJANG PENDIDIKAN</a>
            <a href="#" className="text-gray-700 hover:text-green-600">BERITA</a>
            <a href="#" className="text-gray-700 hover:text-green-600">TENTANG KAMI</a>
            <a href="#" className="text-gray-700 hover:text-green-600">PENDAFTARAN</a>

            {/* Search Bar */}
            <SearchBar value={searchQuery} onChange={setSearchQuery} />
          </nav>

          {/* Mobile menu */}
          <button className="md:hidden">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
      </div>

      {/* Date bar */}
      <div className="bg-gray-100 px-4 py-2 text-center text-sm text-gray-600">
        {getCurrentDate()}
      </div>
    </header>
  );
}
