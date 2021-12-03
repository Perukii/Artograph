
/*
examples/hello_dem_detailed.go
Copyright (C) 2021 Tada Teruki

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation; either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package main

import(
	artograph "../Artograph"
	"fmt"
)

func main(){
	dem := artograph.NewDEM(14)

	// [dem.ElevationAbsM] Absolute value of the maximum/minimum elevation(Meter) (>0.0)
	// (Example : dem.ElevationAbsM = 8000 -> the minimum/maximum elevation : -8000 ~ 8000)
	dem.ElevationAbsM = 8000

	// [dem.UnitKm] The interval of datum point(Km) for generating liver & leveling terrain (>0.0)
	dem.UnitKm = 2

	// [dem.VerticalKm] The vertical width of the DEM data (Km) (>0.0)
	dem.VerticalKm = 1000

	// [dem.HorizontalKm] The horizontal width of the DEM data (Km) (>0.0)
	dem.HorizontalKm = 1000

	// [dem.LevelingIntervalM] The interval of elevation level (Meter) (>0.0)
	dem.LevelingIntervalM = 5

	// [dem.Quality01] The quality of DEM (>0.0) (Recommend : 1.0)
	dem.Quality01 = 1.0

	dem.Generate()

	// ---

	artograph.EnableProcessLog()

	for yKm := 0.0; yKm < dem.HorizontalKm; yKm += dem.HorizontalKm/50.0{
		str := ""
		for xKm := 0.0; xKm < dem.VerticalKm; xKm += dem.HorizontalKm/50.0{

			elevation, _ := dem.GetElevationByKmPoint(xKm, yKm)

			if elevation >= 800{
				str += "@@"
			} else if elevation >= 300 {
				str += "[]"
			} else if elevation >= 0 {
				str += "__"
			} else {
				str += " "
			}

		}
		fmt.Println(str)
	}

}
