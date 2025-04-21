import { useState } from "react";
import Link from "next/link";

export default function LandingPage() {
  const [hoveredCity, setHoveredCity] = useState<string | null>(null);

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 flex flex-col">
      {/* Header */}
      <header className="bg-white shadow-md px-4 py-3 z-10">
        <div className="max-w-6xl mx-auto">
          <h1 className="text-xl font-bold text-blue-600">City Explorer</h1>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 flex flex-col items-center justify-center p-4 md:p-8">
        <div className="text-center mb-12">
          <h2 className="text-3xl md:text-4xl lg:text-5xl font-bold mb-4 text-gray-800">
            Choose Your Destination
          </h2>
          <p className="text-gray-600 max-w-xl mx-auto">
            Select a city to begin your exploration adventure and earn points by
            visiting iconic locations.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8 max-w-4xl w-full">
          {/* Madrid Card - Clickable */}
          <Link href="/madrid-map" className="block">
            <div
              className="relative overflow-hidden rounded-xl shadow-lg transition-all duration-300 transform hover:-translate-y-2 hover:shadow-xl cursor-pointer"
              onMouseEnter={() => setHoveredCity("madrid")}
              onMouseLeave={() => setHoveredCity(null)}
            >
              <div className="h-48 bg-gradient-to-r from-blue-500 to-blue-600 relative">
                <div className="absolute inset-0 flex items-center justify-center">
                  <h3 className="text-5xl md:text-6xl font-bold text-white tracking-wider">
                    MADRID
                  </h3>
                </div>
                <div
                  className={`absolute inset-0 bg-blue-600 transition-opacity duration-300 ${
                    hoveredCity === "madrid" ? "opacity-20" : "opacity-0"
                  }`}
                ></div>
              </div>
              <div className="bg-white p-4">
                <div className="flex justify-between items-center">
                  <div className="flex items-center gap-2">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-5 w-5 text-blue-600"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fillRule="evenodd"
                        d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                        clipRule="evenodd"
                      />
                    </svg>
                    <span className="font-medium text-gray-800">Spain</span>
                  </div>
                  <div className="flex items-center text-blue-600 font-medium">
                    <span>Explore Now</span>
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-5 w-5 ml-1"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fillRule="evenodd"
                        d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </div>
                </div>
                <div className="mt-2 text-sm text-gray-600">
                  Discover 15+ iconic locations and earn points as you explore
                  the beautiful capital of Spain.
                </div>
              </div>
            </div>
          </Link>

          {/* Berlin Card - Non-clickable with Coming Soon */}
          <div
            className="relative overflow-hidden rounded-xl shadow-lg transition-all duration-300 opacity-80"
            onMouseEnter={() => setHoveredCity("berlin")}
            onMouseLeave={() => setHoveredCity(null)}
          >
            <div className="h-48 bg-gradient-to-r from-gray-500 to-gray-600 relative">
              <div className="absolute inset-0 flex items-center justify-center">
                <h3 className="text-5xl md:text-6xl font-bold text-white tracking-wider">
                  BERLIN
                </h3>
              </div>
              <div
                className={`absolute inset-0 bg-gray-600 transition-opacity duration-300 ${
                  hoveredCity === "berlin" ? "opacity-20" : "opacity-0"
                }`}
              ></div>
              {/* Coming Soon Badge */}
              <div className="absolute top-4 right-4 bg-yellow-500 text-white px-3 py-1 rounded-full text-sm font-medium flex items-center gap-1">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-4 w-4"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fillRule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
                    clipRule="evenodd"
                  />
                </svg>
                <span>Coming Soon</span>
              </div>
            </div>
            <div className="bg-white p-4">
              <div className="flex justify-between items-center">
                <div className="flex items-center gap-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 text-gray-500"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fillRule="evenodd"
                      d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                      clipRule="evenodd"
                    />
                  </svg>
                  <span className="font-medium text-gray-800">Germany</span>
                </div>
                <div className="flex items-center text-gray-400 font-medium">
                  <span>Not Available</span>
                </div>
              </div>
              <div className="mt-2 text-sm text-gray-600">
                Berlin exploration is coming soon! Stay tuned for an adventure
                through the historic German capital.
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="bg-white py-4 text-center text-gray-500 text-sm">
        <div className="max-w-6xl mx-auto px-4">
          <p>© 2023 City Explorer. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
}
