package versionnumbers

import "strconv"

func compareVersions(v1, v2 string) int {
	lenV1 := len(v1)
	lenV2 := len(v2)
	i := 0
	j := 0
	startV1 := 0
	startV2 := 0

	for i < lenV1 || j < lenV2 {
		for i < lenV1 && v1[i] != '.' {
			i++
		}

		for j < lenV2 && v2[j] != '.' {
			j++
		}

		var revisionV1 = 0
		var err error = nil
		if i <= lenV1 {
			revisionV1, err = strconv.Atoi(v1[startV1:i])
			if err != nil {
				panic("Conversion error")
			}
			i++
			startV1 = i
		}

		var revisionV2 = 0
		if j <= lenV2 {
			revisionV2, err = strconv.Atoi(v2[startV2:j])
			if err != nil {
				panic("Conversion error")
			}
			j++
			startV2 = j
		}

		if revisionV1 != revisionV2 {
			if revisionV1 > revisionV2 {
				return 1
			}
			return -1
		}
	}

	return 0
}
