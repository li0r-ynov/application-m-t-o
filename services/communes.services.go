package services

import (
 "context"
 "strings"

 "meteo/api"
 "meteo/models"
)

func ListCommunes(nomCommune string) ([]models.Communes, error) {
    return api.SearchCommunesByName(nomCommune)
}

func SuggestCommunes(ctx context.Context, prefix string) ([]models.Communes, error) {
 communes, err := api.GetCommuneSuggestions(ctx, prefix)
 if err != nil {
  return nil, err
 }

 normalizedPrefix := normalizeCommuneName(prefix)
 suggestions := make([]models.Communes, 0, len(communes))
 for _, commune := range communes {
  if strings.HasPrefix(normalizeCommuneName(commune.Name), normalizedPrefix) {
   suggestions = append(suggestions, commune)
  }
 }
 return suggestions, nil
}

func GetCommune(code string) (models.Communes, error) {
 code = strings.TrimSpace(code)
 if code == "" {
  return models.Communes{}, models.ErrBadRequest
 }
 return api.GetCommune(code)
}

func normalizeCommuneName(name string) string {
 name = strings.ToLower(strings.TrimSpace(name))
 replacer := strings.NewReplacer(
  "à", "a", "á", "a", "â", "a", "ä", "a",
  "ç", "c",
  "é", "e", "è", "e", "ê", "e", "ë", "e",
  "î", "i", "ï", "i",
  "ô", "o", "ö", "o",
  "ù", "u", "û", "u", "ü", "u",
  "ÿ", "y", "œ", "oe", "æ", "ae",
 )
 return replacer.Replace(name)
}