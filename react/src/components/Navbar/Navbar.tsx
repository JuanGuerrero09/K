import { Menu, Trophy, User } from "lucide-react";

interface HeaderProps {
  totalPoints: number;
  onToggleSidebar: () => void;
}

export default function Navbar({ totalPoints, onToggleSidebar }: HeaderProps) {
  return (
    <header className="bg-white shadow-md px-4 py-3 z-10 w-full">
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-3">
          <button
            onClick={onToggleSidebar}
            className="p-2 rounded-full hover:bg-gray-100 transition-colors"
          >
            <Menu size={24} />
          </button>
          <h1 className="text-xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
            Madrid Explorer
          </h1>
        </div>

        <div className="flex items-center gap-4">
          <div className="flex items-center gap-2 bg-yellow-50 px-3 py-1.5 rounded-full border border-yellow-200">
            <Trophy size={20} className="text-yellow-500" />
            <span className="font-bold text-yellow-700">{totalPoints} pts</span>
          </div>

          <button className="p-2 rounded-full bg-blue-100 hover:bg-blue-200 transition-colors">
            <User size={20} className="text-blue-700" />
          </button>
        </div>
      </div>
    </header>
  );
}

//   return (
//     <div className="navbar bg-base-100 shadow-sm">
//       <div className="flex-none">
//         <label htmlFor="my-drawer-4" className="drawer-button cursor-pointer">
//           <svg
//             xmlns="http://www.w3.org/2000/svg"
//             fill="none"
//             viewBox="0 0 24 24"
//             className="inline-block h-6 w-6 stroke-current"
//           >
//             <path
//               strokeLinecap="round"
//               strokeLinejoin="round"
//               strokeWidth="2"
//               d="M4 6h16M4 12h16M4 18h16"
//             />
//           </svg>
//         </label>
//       </div>
//       <div className="flex-1">
//         <a className="btn btn-ghost text-xl">Visita Madrid!</a>
//       </div>
//       <div className="flex-none">
//         <button className="btn btn-square btn-ghost">
//           <svg
//             xmlns="http://www.w3.org/2000/svg"
//             fill="none"
//             viewBox="0 0 24 24"
//             className="inline-block h-5 w-5 stroke-current"
//           >
//             {" "}
//             <path
//               strokeLinecap="round"
//               strokeLinejoin="round"
//               strokeWidth="2"
//               d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"
//             ></path>{" "}
//           </svg>
//         </button>
//       </div>
//     </div>
//   );
