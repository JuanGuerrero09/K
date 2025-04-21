import "./../../styles.css";
import "leaflet/dist/leaflet.css";
import { MapContainer, TileLayer } from "react-leaflet";
import MarkerClusterGroup from "react-leaflet-cluster";
import Markers from "../Markers/Markers";
import { Location } from "@/types/location";

interface MapProps {
  locations: Location[];
  selectedLocation: Location | null;
}

export default function Map({ locations, selectedLocation }: MapProps) {
  return (
    <MapContainer
      className=" z-0 border-amber-50"
      center={[40.416775, -3.70379]}
      zoom={13.5}
    >
      {/* OPEN STREEN MAPS TILES */}
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      <MarkerClusterGroup>
        <Markers places={locations} />
      </MarkerClusterGroup>
    </MapContainer>
  );
}
