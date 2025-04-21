export interface Location {
  id: number;
  name: string;
  site: string;
  geoCode: [number, number];
  points: number;
  description: string;
  category: string;
  visited: boolean;
  date_of_completion: string;
  is_unlocked: boolean;
}
