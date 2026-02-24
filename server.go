package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type InviteRequest struct {
	FighterID string `json:"fighterId"`
}
type DrawRequest struct {
	TournamentID string `json:"tournamentId"`
}

type Request struct {
	RequestID    string `json:"request_id"`
	FighterID    string `json:"id_fighter"`
	TournamentID string `json:"tournament_id"`
	Decision     int    `json:"decision"`
	Category     int    `json:"category"`
}

type InvitedParticipant struct {
	FighterID    string `json:"fighter_id"`
	TournamentID string `json:"tournament_id"`
}

type Organazer struct {
	UserID          string `json:"user_id"`
	OrganazerID     string `json:"id"`
	Title           string `json:"title"`
	DescriptionText string `json:"description"`
	FoundationDate  string `json:"founded_date"`
	AddressText     string `json:"address"`
	ContactInfo     string `json:"contact"`
}

type RequestWithTournamentTitle struct {
	RequestID    string `json:"request_id"`
	FighterID    string `json:"id_fighter"`
	Title        string `json:"name"`
	TournamentID string `json:"tournament_id"`
	Category     int    `json:"category"`
	Decision     int    `json:"decision"`
}

type Fighter struct {
	UserID      string  `json:"user_id"`
	ClubID      string  `json:"club_id"`
	FighterID   string  `json:"id"`
	Name        string  `json:"name"`
	Birthday    string  `json:"birthday"`
	Description *string `json:"description"`
	Category    int     `json:"category"`
	Rating      string  `json:"rating"`
	Country     string  `json:"country"`
}

type Club struct {
	UserID          string  `json:"user_id"`
	ClubID          string  `json:"id"`
	Title           string  `json:"title"`
	DescriptionText *string `json:"description"`
	Address         string  `json:"address"`
	Contact         *string `json:"contact"`
	Rating          string  `json:"rating"`
	FoundedDate     string  `json:"founded_date"`
}

type Tournament struct {
	TournamentID string `json:"id"`
	Title        string `json:"name"`
	Address      string `json:"address"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	OrganazerID  string `json:"organazer"`
	RoundsNumber int    `json:"rounds"`
	Category     int    `json:"category"`
}

type TournamentBrackt struct {
	TournamentID string  `json:"tournament_id"`
	Rounds       []Round `json:"rounds"`
}

type Round struct {
	RoundNumber int     `json:"round_number"`
	Matches     []Match `json:"matches"`
}

type Match struct {
	MatchID             string `json:"match_id"`
	TournamentID        string `json:"tournament_id"`
	MatchNumber         int    `json:"match_number"`
	RoundNumber         int    `json:"round_number"`
	TopParticipantID    string `json:"top_participant_id"`
	BottomParticipantID string `json:"bottom_participant_id"`
	WhenPlayed          string `json:"when_played"`
	WinnerID            string `json:"winner_id"`
	Top_score           int    `json:"top_score"`
	Bottom_score        int    `json:"bottom_score"`
}

type SelectedMember struct {
	RequestID string `json:"REQUEST_ID"`
	Win       int    `json:"WIN"`
	Loss      int    `json:"LOSS"`
}

type MatchScoreUpdate struct {
	MatchID      string `json:"match_id"`
	Top_score    int    `json:"top_score"`
	Bottom_score int    `json:"bottom_score"`
}
type Participant struct {
	ParticipantID string `json:"participantId"`
	Name          string `json:"name"`
}

type TableResults struct {
	TournamentID  string `json:"tournament_id"`
	RoundNumber   int    `json:"roundNumber"`
	MatchID       string `json:"match_id"`
	ParticipantID string `json:"participant_id"`
	Place         string `json:"place"`
}

type UserProfile struct {
	UserProfileID string `json:"user_profile_id"`
	Username      string `json:"username"`
	ProfileType   int    `json:"profile_type"`
	Email         string `json:"e_mail"`
	Password      string `json:"password"`
}
type UserProfile2 struct {
	UserProfileID string `json:"user_profile_id"`
	Username      string `json:"username"`
	ProfileType   int    `json:"profile_type"`
	Linked_id_1   string `json:"linked_id_1"`
	Linked_id_2   string `json:"linked_id_2"`
	Linked_id_3   string `json:"linked_id_3"`
	Email         string `json:"e_mail"`
	Password      string `json:"password"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var db *sql.DB

func main() {
	var err error

	server := "localhost"
	port := 1433
	user := "dbuser"
	password := "1234"
	database := "DB"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных: ", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка при проверке соединения с базой данных: ", err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/clubs", corsHandler(getClubs)).Methods("GET")
	r.HandleFunc("/club-profile/{id}", corsHandler(getClub)).Methods("GET")
	r.HandleFunc("/club-profile/{id}/members", corsHandler(getClubMembers)).Methods("GET")
	r.HandleFunc("/club-profile/{id}/update", corsHandler(updateClub)).Methods("POST")
	r.HandleFunc("/club-profile/{id}/delete", corsHandler(deleteClub)).Methods("GET")
	r.HandleFunc("/club-profile/{id}/addMember/{fighter_id}", corsHandler(addMemberToClub)).Methods("GET")
	r.HandleFunc("/club-profile/{id}/kickMember/{fighter_id}", corsHandler(kickFromClub)).Methods("GET")

	r.HandleFunc("/add_club/{user_id}", corsHandler(addClub)).Methods("POST")
	r.HandleFunc("/add_fighter/{user_id}", corsHandler(addFighter)).Methods("POST")
	r.HandleFunc("/fighters", corsHandler(getFighters)).Methods("GET")
	r.HandleFunc("/fighters_no_club", corsHandler(getFightersWithOutClub)).Methods("GET")
	r.HandleFunc("/fighter/{id}/delete", corsHandler(deleteFighter)).Methods("GET")
	r.HandleFunc("/fighter/{id}/update", corsHandler(updateFighter)).Methods("POST")
	r.HandleFunc("/fighter/{id}/requests", corsHandler(getRequestsByFighterId)).Methods("GET")

	r.HandleFunc("/fighter/{id}", corsHandler(getFighter)).Methods("GET")
	r.HandleFunc("/sendRequest", corsHandler(sendRequest)).Methods("POST")
	r.HandleFunc("/tournaments", corsHandler(getTournaments)).Methods("GET")
	r.HandleFunc("/tournament/{id}", corsHandler(getTournament)).Methods("GET")
	r.HandleFunc("/tournament/{id}/invite", corsHandler(inviteRequest)).Methods("POST")
	r.HandleFunc("/tournament/{tournamentId}/delete_invite/{inviteId}", corsHandler(deleteinviteRequest)).Methods("GET")
	r.HandleFunc("/tournament/{tournamentId}/delete", corsHandler(deleteTournament)).Methods("GET")
	r.HandleFunc("/tournament/{tournamentId}/update", corsHandler(updateTournament)).Methods("POST")

	r.HandleFunc("/add_organazer/{user_id}", corsHandler(addOrganazer)).Methods("POST")
	r.HandleFunc("/tournament/{id}/bracket", corsHandler(getTournamentBracket)).Methods("GET")
	r.HandleFunc("/organazers", corsHandler(getOrganazers)).Methods("GET")
	r.HandleFunc("/organazer-profile/{id}", corsHandler(getOrganazer)).Methods("GET")
	r.HandleFunc("/organazer-profile/{id}/tournaments_organazed", corsHandler(getOrganazedTournaments)).Methods("GET")
	r.HandleFunc("/organazer/create_tournament", corsHandler(addTournament)).Methods("POST")
	r.HandleFunc("/requests/{tournamentId}", corsHandler(getRequests)).Methods("GET")
	r.HandleFunc("/invited/{tournamentId}", corsHandler(getInvited)).Methods("GET")
	r.HandleFunc("/approve/{requestId}", corsHandler(approveRequest)).Methods("POST")
	r.HandleFunc("/deny/{requestId}", corsHandler(denyRequest)).Methods("POST")

	r.HandleFunc("/tournament/{tournamentId}/selected-members", corsHandler(getParticipantsNamed)).Methods("GET")
	r.HandleFunc("/tournament/{tournamentId}/draw", corsHandler(drawMatchesHandler)).Methods("GET")
	r.HandleFunc("/tournament/{tournamentId}/conduct", corsHandler(conductTournament)).Methods("GET")
	r.HandleFunc("/tournament/{tournamentId}/update-match-score", corsHandler(updateMatchScoreHandler)).Methods("POST")
	r.HandleFunc("/tournament/{tournament_id}/matches", corsHandler(getMatchesById)).Methods("GET")
	r.HandleFunc("/tournament/{tournament_id}/results", corsHandler(getResultsHandler)).Methods("GET")

	r.HandleFunc("/requests/{requestId}/delete", corsHandler(deleteRequest)).Methods("GET")
	r.HandleFunc("/organazer-profile/{id}/update", corsHandler(updateOrganazer)).Methods("POST")
	r.HandleFunc("/organazer-profile/{id}/delete", corsHandler(deleteOrganazer)).Methods("GET")

	r.HandleFunc("/registration", corsHandler(registerUser)).Methods("POST")
	r.HandleFunc("/login", corsHandler(loginUser)).Methods("POST")
	r.HandleFunc("/user-profile/{id}", corsHandler(getUserProfile)).Methods("GET")
	r.HandleFunc("/user-profile/{e_mail}/email", corsHandler(getUserByEmail)).Methods("GET")
	r.HandleFunc("/user-profile/{username}/username", corsHandler(getUserByUsername)).Methods("GET")
	r.HandleFunc("/user-profile/{id}/update", corsHandler(updateUserProfile)).Methods("POST")
	r.HandleFunc("/user-profile/{id}/delete", corsHandler(deleteUserProfile)).Methods("GET")

	log.Println("Сервер запущен на порту 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}

func writeToFile(data []byte) error {
	file, err := os.Create("clubs_data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	log.Println("Данные успешно записаны в файл clubs_data.json")
	return nil
}

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)

			return
		}
		h(w, r)
	}
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-Custom-Header, my_header")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getClubs(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT convert(nvarchar(50), CLUB_ID), TITLE, DESCRIPTION_TEXT, FOUNDED_DATE, CONTACT, RATING, COUNTRY FROM CLUB")
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clubs []Club
	for rows.Next() {
		var club Club
		if err := rows.Scan(&club.ClubID, &club.Title, &club.DescriptionText, &club.FoundedDate, &club.Contact, &club.Rating, &club.Address); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}
		clubs = append(clubs, club)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"clublist": clubs,
	}

	clubsJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(clubsJSON)
}

func getClub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var club Club
	err := db.QueryRow("SELECT convert(nvarchar(50), CLUB_ID), TITLE, DESCRIPTION_TEXT, FOUNDED_DATE, CONTACT, convert(nvarchar(50), RATING), COUNTRY FROM CLUB WHERE CLUB_ID = @p1", id).
		Scan(&club.ClubID, &club.Title, &club.DescriptionText, &club.FoundedDate, &club.Contact, &club.Rating, &club.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Клуб не найден", http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	clubJSON, err := json.Marshal(club)
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(clubJSON)
}

func getClubMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	rows, err := db.Query(`SELECT CONVERT(NVARCHAR(50),FIGHTER.CLUB_ID), CONVERT(NVARCHAR(50),FIGHTER.FIGHTER_ID), FIGHTER.NAME_TEXT, FIGHTER.BIRTHDAY,
        FIGHTER.DESCRIPTION_TEXT, FIGHTER.CATEGORY, convert(nvarchar(50), FIGHTER.RATING), FIGHTER.COUNTRY 
        FROM FIGHTER 
        INNER JOIN CLUB ON CLUB.CLUB_ID = FIGHTER.CLUB_ID 
        WHERE CLUB.CLUB_ID = @p1`, id)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var fighters []Fighter
	for rows.Next() {
		var fighter Fighter
		if err := rows.Scan(&fighter.ClubID, &fighter.FighterID, &fighter.Name, &fighter.Birthday, &fighter.Description, &fighter.Category, &fighter.Rating, &fighter.Country); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fighters = append(fighters, fighter)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"memberlist": fighters,
	}

	fightersJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fightersJSON)
}

func addClub(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	user_id := vars["user_id"]

	var club Club
	err := json.NewDecoder(r.Body).Decode(&club)
	log.Print(club.Contact, club.DescriptionText)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON данных: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Print(club.Contact, club.DescriptionText)

	if club.Title == "" || club.FoundedDate == "" || club.Rating == "" || club.Address == "" {
		http.Error(w, "Не все обязательные поля заполнены", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", club.FoundedDate); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	clubID := uuid.New()
	clubIDString := clubID.String()

	ratingFloat, err := strconv.ParseFloat(club.Rating, 32)
	if err != nil {
		http.Error(w, "Неверный формат рейтинга: "+err.Error(), http.StatusBadRequest)
		return
	}

	tsql := fmt.Sprintf("INSERT INTO CLUB (USER_PROFILE_ID, CLUB_ID, TITLE, DESCRIPTION_TEXT, FOUNDED_DATE, CONTACT, RATING, COUNTRY) VALUES ('%s','%s', '%s', '%s', convert(date, '%s'), '%s', %f, '%s');",
		user_id, clubIDString, club.Title, *club.DescriptionText, club.FoundedDate, *club.Contact, ratingFloat, club.Address)

	_, err = db.Exec(tsql)
	if err != nil {
		log.Print(err)
		http.Error(w, "Ошибка при выполнении запроса к базе данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func addMemberToClub(w http.ResponseWriter, r *http.Request) {
	clubID := mux.Vars(r)["id"]
	fighterID := mux.Vars(r)["fighter_id"]

	if clubID == "" || fighterID == "" {
		http.Error(w, "ID клуба или ID бойца не указаны", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("UPDATE FIGHTER SET CLUB_ID = '%s' WHERE FIGHTER_ID = '%s'", clubID, fighterID)

	_, err := db.Exec(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при добавлении участника в клуб: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func kickFromClub(w http.ResponseWriter, r *http.Request) {
	clubID := mux.Vars(r)["id"]
	fighterID := mux.Vars(r)["fighter_id"]

	if clubID == "" || fighterID == "" {
		http.Error(w, "ID клуба или ID бойца не указаны", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("UPDATE FIGHTER SET CLUB_ID = NULL WHERE FIGHTER_ID = '%s'", fighterID)

	_, err := db.Exec(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при добавлении участника в клуб: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func addFighter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	user_id := vars["user_id"]

	var fighter Fighter
	err := json.NewDecoder(r.Body).Decode(&fighter)
	log.Print(fighter.Name, fighter.Country, fighter.Description)

	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON данных: "+err.Error(), http.StatusBadRequest)
		return
	}

	if fighter.Name == "" || fighter.Birthday == "" || fighter.Category == 0 || fighter.Rating == "" || fighter.Country == "" {
		http.Error(w, "Не все обязательные поля заполнены", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", fighter.Birthday); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	fighterId := uuid.New()
	fighterID := fighterId.String()

	ratingFloat, err := strconv.ParseFloat(fighter.Rating, 32)
	if err != nil {
		http.Error(w, "Неверный формат рейтинга: "+err.Error(), http.StatusBadRequest)
		return
	}

	tsqlFighter := fmt.Sprintf("INSERT INTO FIGHTER (USER_PROFILE_ID, FIGHTER_ID, NAME_TEXT, BIRTHDAY, DESCRIPTION_TEXT, CATEGORY, RATING, COUNTRY) VALUES ('%s','%s', '%s', convert(date, '%s'), '%s', %d, %f, '%s');",
		user_id, fighterID, fighter.Name, fighter.Birthday, *fighter.Description, fighter.Category, ratingFloat, fighter.Country)

	_, err = db.Exec(tsqlFighter)
	if err != nil {
		log.Print(err)
		http.Error(w, "Ошибка при выполнении запроса к базе данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getFighter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var fighter Fighter

	query1 := fmt.Sprintf(`
        SELECT CONVERT(nvarchar(50), CLUB_ID), CONVERT(nvarchar(50), FIGHTER_ID), NAME_TEXT, BIRTHDAY, DESCRIPTION_TEXT, CATEGORY, CONVERT(nvarchar(50), RATING), COUNTRY 
        FROM FIGHTER 
        WHERE FIGHTER_ID = '%s' AND CLUB_ID IS NOT NULL`, id)

	err := db.QueryRow(query1).Scan(&fighter.ClubID, &fighter.FighterID, &fighter.Name, &fighter.Birthday, &fighter.Description, &fighter.Category, &fighter.Rating, &fighter.Country)
	if err != nil {
		if err == sql.ErrNoRows {
			query2 := fmt.Sprintf(`
                SELECT CONVERT(nvarchar(50), FIGHTER_ID), NAME_TEXT, BIRTHDAY, DESCRIPTION_TEXT, CATEGORY, CONVERT(nvarchar(50), RATING), COUNTRY 
                FROM FIGHTER 
                WHERE FIGHTER_ID = '%s' AND CLUB_ID IS NULL`, id)

			err2 := db.QueryRow(query2).Scan(&fighter.FighterID, &fighter.Name, &fighter.Birthday, &fighter.Description, &fighter.Category, &fighter.Rating, &fighter.Country)
			if err2 != nil {
				log.Println(err2)
				http.Error(w, "Fighter not found", http.StatusNotFound)
				return
			}
		} else {
			log.Println(err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	}

	fighterJSON, err := json.Marshal(fighter)
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fighterJSON)
}

func getFighters(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT CONVERT(NVARCHAR(50), FIGHTER_ID), NAME_TEXT, BIRTHDAY, DESCRIPTION_TEXT, CATEGORY, CONVERT(NVARCHAR(50), RATING), COUNTRY, CONVERT(NVARCHAR(50), CLUB_ID) FROM FIGHTER where CLUB_ID IS NOT NULL")
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var fighters []Fighter
	for rows.Next() {
		var fighter Fighter
		var descText sql.NullString

		if err := rows.Scan(&fighter.FighterID, &fighter.Name, &fighter.Birthday, &descText, &fighter.Category, &fighter.Rating, &fighter.Country, &fighter.ClubID); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if descText.Valid {
			fighter.Description = &descText.String
		} else {
			fighter.Description = nil
		}

		fighters = append(fighters, fighter)
	}

	rows, err = db.Query("SELECT CONVERT(NVARCHAR(50), FIGHTER_ID), NAME_TEXT, BIRTHDAY, DESCRIPTION_TEXT, CATEGORY, CONVERT(NVARCHAR(50), RATING), COUNTRY FROM FIGHTER where CLUB_ID IS NULL")
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса 2: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var fighter Fighter
		var descText sql.NullString

		if err := rows.Scan(&fighter.FighterID, &fighter.Name, &fighter.Birthday, &descText, &fighter.Category, &fighter.Rating, &fighter.Country); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if descText.Valid {
			fighter.Description = &descText.String
		} else {
			fighter.Description = nil
		}
		fighters = append(fighters, fighter)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"fighterlist": fighters,
	}

	fightersJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fightersJSON)
}

func getFightersWithOutClub(w http.ResponseWriter, r *http.Request) {

	var fighters []Fighter

	rows, err := db.Query("SELECT CONVERT(NVARCHAR(50), FIGHTER_ID), NAME_TEXT, BIRTHDAY, DESCRIPTION_TEXT, CATEGORY, CONVERT(NVARCHAR(50), RATING), COUNTRY FROM FIGHTER where CLUB_ID IS NULL")
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса 2: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var fighter Fighter
		var descText sql.NullString

		if err := rows.Scan(&fighter.FighterID, &fighter.Name, &fighter.Birthday, &descText, &fighter.Category, &fighter.Rating, &fighter.Country); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if descText.Valid {
			fighter.Description = &descText.String
		} else {
			fighter.Description = nil
		}
		fighters = append(fighters, fighter)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"fighterlist": fighters,
	}

	fightersJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fightersJSON)
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func addTournament(w http.ResponseWriter, r *http.Request) {
	var req Tournament

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println(req.OrganazerID)

	query := `
    INSERT INTO TOURNAMENT (TITLE, START_DATE, END_DATE, ORGANAZER_ID, ROUNDS_NUMBER, CATEGORY, ADDRESS)
    VALUES (@Title, @StartDate, @EndDate, @OrganazerID, @RoundsNumber, @Category, @Address)
    `
	_, err = db.Exec(query,
		sql.Named("Title", req.Title),
		sql.Named("StartDate", req.StartDate),
		sql.Named("EndDate", req.EndDate),
		sql.Named("OrganazerID", req.OrganazerID),
		sql.Named("RoundsNumber", req.RoundsNumber),
		sql.Named("Category", req.Category),
		sql.Named("Address", req.Address),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c := 0
	totalRounds := req.RoundsNumber

	for roundNumber := 1; roundNumber <= totalRounds; roundNumber++ {
		matchesInRound := int(math.Pow(2, float64(totalRounds-roundNumber)))

		for matchNumber := 1; matchNumber <= matchesInRound; matchNumber++ {
			c++

			queryMatch := `
            INSERT INTO MATCH_TABLE (TOURNAMENT_ID, ROUND_NUMBER, MATCH_NUMBER, TOP_PARTICIPANT_ID, BOTTOM_PARTICIPANT_ID, WHEN_PLAYED, WINNER_ID, TOP_SCORE, BOTTOM_SCORE)
            VALUES ((SELECT TOURNAMENT_ID FROM TOURNAMENT WHERE TITLE = @Title), @RoundNumber, @MatchNumber, NULL, NULL, NULL, NULL, NULL, NULL)
            `
			_, err := db.Exec(queryMatch,
				sql.Named("Title", req.Title),
				sql.Named("RoundNumber", roundNumber),
				sql.Named("MatchNumber", c),
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func getTournaments(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT 
                            CONVERT(NVARCHAR(100), TOURNAMENT_ID), 
                            TITLE, 
                            START_DATE, 
                            END_DATE, 
                            CONVERT(NVARCHAR(100), Organazer_ID), 
                            ROUNDS_NUMBER,
							CATEGORY,
							ADDRESS
                          FROM 
                            TOURNAMENT`)
	if err != nil {
		log.Print(err)
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tournaments []Tournament
	for rows.Next() {
		var tournament Tournament

		if err := rows.Scan(&tournament.TournamentID, &tournament.Title, &tournament.StartDate, &tournament.EndDate, &tournament.OrganazerID, &tournament.RoundsNumber, &tournament.Category, &tournament.Address); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}

		tournaments = append(tournaments, tournament)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"tournamentlist": tournaments,
	}

	tournamentsJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tournamentsJSON)
}

func getTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tournament Tournament
	err := db.QueryRow("SELECT convert(nvarchar(50), TOURNAMENT_ID), TITLE, START_DATE, END_DATE, convert(nvarchar(50), Organazer_ID), ROUNDS_NUMBER, CATEGORY, ADDRESS FROM TOURNAMENT WHERE TOURNAMENT_ID = @p1", id).
		Scan(&tournament.TournamentID, &tournament.Title, &tournament.StartDate, &tournament.EndDate, &tournament.OrganazerID, &tournament.RoundsNumber, &tournament.Category, &tournament.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Турнир не найден", http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	tournamentJSON, err := json.MarshalIndent(tournament, "", " ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tournamentJSON)
}

func getTournamentBracket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tournament, err := getTournamentFromDB(db, id)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tournamentJSON, err := json.MarshalIndent(tournament, "", " ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tournamentJSON)
}

func getTournamentFromDB(db *sql.DB, tournamentID string) (TournamentBrackt, error) {
	var tournament TournamentBrackt
	tournament.TournamentID = tournamentID

	queryMatches := fmt.Sprintf(`SELECT convert(nvarchar(50), MATCH_ID), ROUND_NUMBER, convert(nvarchar(50),TOP_PARTICIPANT_ID), convert(nvarchar(50),BOTTOM_PARTICIPANT_ID),
        WHEN_PLAYED, convert(nvarchar(50),WINNER_ID), TOP_SCORE, BOTTOM_SCORE 
        FROM MATCH_TABLE WHERE TOURNAMENT_ID = '%s' ORDER BY ROUND_NUMBER, MATCH_ID`, tournamentID)
	rowsMatches, err := db.Query(queryMatches)
	if err != nil {
		return tournament, fmt.Errorf("ошибка при извлечении данных матчей: %v", err)
	}
	defer rowsMatches.Close()

	roundsMap := make(map[int][]Match)

	for rowsMatches.Next() {
		var match Match
		var top, bottom, whenPlayed, winner sql.NullString
		var topScore, bottomScore sql.NullInt16

		if err := rowsMatches.Scan(&match.MatchID, &match.RoundNumber, &top, &bottom, &whenPlayed, &winner, &topScore, &bottomScore); err != nil {
			return tournament, fmt.Errorf("ошибка при сканировании матчей: %v", err)
		}

		if top.Valid {
			match.TopParticipantID = top.String
			fmt.Print(match.TopParticipantID)
		}
		if bottom.Valid {
			match.BottomParticipantID = bottom.String
		}
		if whenPlayed.Valid {
			match.WhenPlayed = whenPlayed.String
		}
		if winner.Valid {
			match.WinnerID = winner.String
		}
		if topScore.Valid {
			match.Top_score = int(topScore.Int16)
		}
		if bottomScore.Valid {
			match.Bottom_score = int(bottomScore.Int16)
		}

		roundsMap[match.RoundNumber] = append(roundsMap[match.RoundNumber], match)
	}

	for roundNumber, matches := range roundsMap {
		round := Round{
			RoundNumber: roundNumber,
			Matches:     matches,
		}
		tournament.Rounds = append(tournament.Rounds, round)
	}

	return tournament, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getOrganazers(w http.ResponseWriter, r *http.Request) {

	query := "SELECT convert(nvarchar(50), ORGANAZER_ID), TITLE, DESCRIPTION_TEXT, FOUNDATION_DATE, ADDRESS_TEXT, CONTACT_INFO FROM ORGANAZER_TOUR"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var organazers []Organazer
	for rows.Next() {
		var organazer Organazer
		var descriptionText sql.NullString
		var addressText sql.NullString
		var contactInfo sql.NullString

		if err := rows.Scan(&organazer.OrganazerID, &organazer.Title, &descriptionText, &organazer.FoundationDate, &addressText, &contactInfo); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if descriptionText.Valid {
			organazer.DescriptionText = descriptionText.String
		}
		if addressText.Valid {
			organazer.AddressText = addressText.String
		}
		if contactInfo.Valid {
			organazer.ContactInfo = contactInfo.String
		}
		organazers = append(organazers, organazer)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"Organazerlist": organazers,
	}

	OrganazersJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(OrganazersJSON)
}

func getOrganazer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var organazer Organazer
	var descriptionText sql.NullString
	var addressText sql.NullString
	var contactInfo sql.NullString
	err := db.QueryRow("SELECT convert(nvarchar(50), ORGANAZER_ID), TITLE, DESCRIPTION_TEXT, FOUNDATION_DATE, ADDRESS_TEXT, CONTACT_INFO FROM ORGANAZER_TOUR WHERE ORGANAZER_ID = @p1", id).
		Scan(&organazer.OrganazerID, &organazer.Title, &descriptionText, &organazer.FoundationDate, &addressText, &contactInfo)
	if descriptionText.Valid {
		organazer.DescriptionText = descriptionText.String
	}
	if addressText.Valid {
		organazer.AddressText = addressText.String
	}
	if contactInfo.Valid {
		organazer.ContactInfo = contactInfo.String
	}
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Организатор не найден", http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	organizerJSON, err := json.Marshal(organazer)
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(organizerJSON)
}

func getOrganazedTournaments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	query := fmt.Sprintf("SELECT CONVERT(NVARCHAR(100), TOURNAMENT.TOURNAMENT_ID), TOURNAMENT.TITLE, TOURNAMENT.START_DATE, TOURNAMENT.END_DATE, CONVERT(NVARCHAR(100), TOURNAMENT.ORGANAZER_ID), TOURNAMENT.ROUNDS_NUMBER, TOURNAMENT.CATEGORY, TOURNAMENT.ADDRESS FROM TOURNAMENT WHERE TOURNAMENT.ORGANAZER_ID = '%s';", id)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tournaments []Tournament
	for rows.Next() {
		var tournament Tournament

		if err := rows.Scan(&tournament.TournamentID, &tournament.Title, &tournament.StartDate, &tournament.EndDate, &tournament.OrganazerID, &tournament.RoundsNumber, &tournament.Category, &tournament.Address); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}

		tournaments = append(tournaments, tournament)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"tournamentlistorg": tournaments,
	}

	tournamentsJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tournamentsJSON)
}

func addOrganazer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	user_id := vars["user_id"]

	var organazer Organazer
	err := json.NewDecoder(r.Body).Decode(&organazer)
	log.Print(organazer.Title)
	log.Print(organazer.AddressText)
	log.Print(organazer.DescriptionText)
	log.Print(organazer.FoundationDate)

	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON данных: "+err.Error(), http.StatusBadRequest)
		return
	}

	if organazer.Title == "" || organazer.FoundationDate == "" || organazer.AddressText == "" {
		http.Error(w, "Не все обязательные поля заполнены", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", organazer.FoundationDate); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	organizerID := uuid.New()
	organizerIDString := organizerID.String()

	query := `
        INSERT INTO ORGANAZER_TOUR (USER_PROFILE_ID, ORGANAZER_ID, TITLE, DESCRIPTION_TEXT, FOUNDATION_DATE, ADDRESS_TEXT, CONTACT_INFO)
        VALUES (@User_id, @OrganizerID, @Title, @DescriptionText, @FoundationDate, @AddressText, @ContactInfo)`

	_, err = db.Exec(query,
		sql.Named("User_id", user_id),
		sql.Named("OrganizerID", organizerIDString),
		sql.Named("Title", organazer.Title),
		sql.Named("DescriptionText", organazer.DescriptionText),
		sql.Named("FoundationDate", organazer.FoundationDate),
		sql.Named("AddressText", organazer.AddressText),
		sql.Named("ContactInfo", organazer.ContactInfo),
	)
	if err != nil {
		log.Print(err)
		http.Error(w, "Ошибка при выполнении запроса к базе данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func sendRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%s", string(req.FighterID))

	req.Decision = 0

	query := `
        INSERT INTO REQUEST (FIGHTER_ID, TOURNAMENT_ID, DECISION)
        VALUES (@p1, @p2, @p3)`

	_, err = db.Exec(query, req.FighterID, req.TournamentID, req.Decision)
	if err != nil {
		http.Error(w, "Error inserting record into database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func inviteRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	tournamentID := vars["id"]

	if tournamentID == "" {
		http.Error(w, "Missing tournament ID", http.StatusBadRequest)
		return
	}

	var invite InviteRequest

	err := json.NewDecoder(r.Body).Decode(&invite)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if invite.FighterID == "" || tournamentID == "" {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO INVITED_PARTICIPANTS (FIGHTER_ID, TOURNAMENT_ID) VALUES (@fighterID, @tournamentID)"

	_, err = db.Exec(query, sql.Named("fighterID", invite.FighterID), sql.Named("tournamentID", tournamentID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing query: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentId := vars["tournamentId"]
	query := fmt.Sprintf("SELECT convert(nvarchar(50), REQUEST_ID), convert(nvarchar(50),FIGHTER_ID), convert(nvarchar(50),TOURNAMENT_ID), DECISION FROM REQUEST WHERE TOURNAMENT_ID = '%s'", tournamentId)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var requests []Request
	for rows.Next() {
		var req Request
		if err := rows.Scan(&req.RequestID, &req.FighterID, &req.TournamentID, &req.Decision); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requests = append(requests, req)
	}

	response := map[string]interface{}{
		"requests": requests,
	}

	requestsJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%s", string(requestsJSON))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(requestsJSON)
}

func getInvited(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentId := vars["tournamentId"]
	query := fmt.Sprintf("SELECT convert(nvarchar(50),FIGHTER_ID), convert(nvarchar(50),TOURNAMENT_ID) FROM INVITED_PARTICIPANTS WHERE TOURNAMENT_ID = '%s'", tournamentId)

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var invitedParticipants []InvitedParticipant
	for rows.Next() {
		var participant InvitedParticipant
		if err := rows.Scan(&participant.FighterID, &participant.TournamentID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		invitedParticipants = append(invitedParticipants, participant)
	}
	response := map[string]interface{}{
		"invitedParticipants": invitedParticipants,
	}

	invitedParticipantsJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(invitedParticipantsJSON)
}

func approveRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	query := fmt.Sprintf("UPDATE REQUEST SET DECISION = 1 WHERE REQUEST_ID = '%s'", requestId)
	_, err := db.Exec(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func denyRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	query := fmt.Sprintf("UPDATE REQUEST SET DECISION = 0 WHERE REQUEST_ID = '%s'", requestId)

	_, err := db.Exec(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func conductTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentID := vars["tournamentId"]

	var roundsNumber int
	queryRounds := fmt.Sprintf(`SELECT ROUNDS_NUMBER FROM TOURNAMENT WHERE TOURNAMENT_ID = '%s';`, tournamentID)

	err := db.QueryRow(queryRounds).Scan(&roundsNumber)
	if err != nil {
		log.Println("Ошибка при получения числа раундов:", err)
		http.Error(w, "Ошибка при получении информации о турнире", http.StatusInternalServerError)
		return
	}

	maxParticipants := 1 << roundsNumber

	var approvedCount int
	queryApprovedCount := fmt.Sprintf(`SELECT COUNT(*) FROM REQUEST 
WHERE DECISION = 1 AND TOURNAMENT_ID = '%s';`, tournamentID)

	err = db.QueryRow(queryApprovedCount).Scan(&approvedCount)
	if err != nil {
		log.Println("Ошибка при подсчете подтвержденных участников:", err)
		http.Error(w, "Ошибка при подсчете участников", http.StatusInternalServerError)
		return
	}

	if approvedCount < maxParticipants {
		http.Error(w, fmt.Sprintf("Недостаточное количество участников: ожидается %d, фактически %d", maxParticipants, approvedCount), http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO SELECTED_MEMBER(REQUEST_ID, FIGHTER_ID, TOURNAMENT_ID)
SELECT R.REQUEST_ID, F.FIGHTER_ID, T.TOURNAMENT_ID
FROM REQUEST AS R 
INNER JOIN FIGHTER AS F ON R.FIGHTER_ID = F.FIGHTER_ID 
INNER JOIN TOURNAMENT AS T ON R.TOURNAMENT_ID = T.TOURNAMENT_ID
WHERE R.DECISION = 1 AND R.INVITED_PARTICIPANTS_ID IS NOT NULL AND R.TOURNAMENT_ID = '%s';`, tournamentID)

	query2 := fmt.Sprintf(`INSERT INTO SELECTED_MEMBER(REQUEST_ID, FIGHTER_ID, TOURNAMENT_ID)
SELECT R.REQUEST_ID, F.FIGHTER_ID, T.TOURNAMENT_ID
FROM REQUEST AS R 
INNER JOIN FIGHTER AS F ON R.FIGHTER_ID = F.FIGHTER_ID 
INNER JOIN TOURNAMENT AS T ON R.TOURNAMENT_ID = T.TOURNAMENT_ID
WHERE R.DECISION = 1 AND R.INVITED_PARTICIPANTS_ID IS NULL AND R.TOURNAMENT_ID = '%s';`, tournamentID)

	_, err = db.Exec(query)
	if err != nil {
		log.Println("Ошибка при добавлении участников:", err)
		http.Error(w, "Ошибка при добавлении участников", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(query2)
	if err != nil {
		log.Println("Ошибка при добавлении участников:", err)
		http.Error(w, "Ошибка при добавлении участников", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getParticipants(tournamentID string, roundNumber int) ([]string, error) {
	participants := []string{}

	rows, err := db.Query(`
    SELECT CONVERT(nvarchar(50), SELECTED_MEMBER.PARTICIPANT_ID) 
    FROM SELECTED_MEMBER 
    INNER JOIN TOURNAMENT 
    ON SELECTED_MEMBER.TOURNAMENT_ID = TOURNAMENT.TOURNAMENT_ID 
    AND TOURNAMENT.TOURNAMENT_ID = @p1`, tournamentID)
	if err != nil {
		fmt.Print(3)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var participantId string
		if err := rows.Scan(&participantId); err != nil {
			return nil, err
		}
		participants = append(participants, participantId)
	}

	return participants, nil
}

func nextMatchByTree(totalRounds, currentMatch int) int {
	levelStart := 1
	for i := totalRounds; i > 0; i-- {
		levelEnd := levelStart + int(math.Pow(2, float64(i-1))) - 1
		if currentMatch >= levelStart && currentMatch <= levelEnd {
			matchPosition := currentMatch - levelStart
			nextLevelStart := levelEnd + 1
			nextMatch := nextLevelStart + matchPosition/2
			if nextMatch >= levelEnd+1 {
				return nextMatch
			}
			return -1
		}
		levelStart = levelEnd + 1
	}
	return -1
}

func assignNextMatch(winnerID string, matchID string, w http.ResponseWriter) error {
	fmt.Printf("Winner ID: %s, Match ID: %s\n", winnerID, matchID)
	row := db.QueryRow(fmt.Sprintf("SELECT MATCH_NUMBER, ROUND_NUMBER, CONVERT(NVARCHAR(50), TOURNAMENT_ID) FROM MATCH_TABLE WHERE MATCH_ID = '%s'", matchID))
	var tournamentID string
	fmt.Print("\namogus\n")
	var currentMatchNumber, currentRoundNumber int
	if err := row.Scan(&currentMatchNumber, &currentRoundNumber, &tournamentID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return fmt.Errorf("match with ID %s not found", matchID)
		}
		return err
	}
	fmt.Print("amogus2")
	var maxRounds int
	err := db.QueryRow("SELECT ROUNDS_NUMBER FROM TOURNAMENT WHERE TOURNAMENT.TOURNAMENT_ID = @p1", sql.Named("p1", tournamentID)).Scan(&maxRounds)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return err
	}
	fmt.Print("\nтекущий матч и раунд", currentMatchNumber, " ", currentRoundNumber, "\n")
	nextMatchNumber := nextMatchByTree(maxRounds, currentMatchNumber)
	fmt.Print("\nследующий", nextMatchNumber, "\n")

	nextRoundNumber := currentRoundNumber + 1
	fmt.Print("\nследующий раунд", nextRoundNumber, "\n")
	var position string
	if currentMatchNumber%2 == 1 {
		position = "bottom"
	} else {
		position = "top"
	}

	var topParticipantID, bottomParticipantID sql.NullString
	query := fmt.Sprintf("SELECT convert(nvarchar(50), TOP_PARTICIPANT_ID), convert(nvarchar(50), BOTTOM_PARTICIPANT_ID) FROM MATCH_TABLE WHERE MATCH_NUMBER = %d AND ROUND_NUMBER = %d", nextMatchNumber, nextRoundNumber)

	err = db.QueryRow(query).Scan(&topParticipantID, &bottomParticipantID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if position == "top" {
		updateTopParticipant := fmt.Sprintf("UPDATE MATCH_TABLE SET TOP_PARTICIPANT_ID = '%s' WHERE MATCH_NUMBER = %d AND ROUND_NUMBER = %d", winnerID, nextMatchNumber, nextRoundNumber)
		fmt.Print(updateTopParticipant, "\ntop")
		_, err := db.Exec(updateTopParticipant)
		return err
	} else {
		updateBottomParticipant := fmt.Sprintf("UPDATE MATCH_TABLE SET BOTTOM_PARTICIPANT_ID = '%s' WHERE MATCH_NUMBER = %d AND ROUND_NUMBER = %d", winnerID, nextMatchNumber, nextRoundNumber)
		fmt.Print(updateBottomParticipant, "\nbottom")

		_, err := db.Exec(updateBottomParticipant)
		return err
	}
}

func drawMatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentId := vars["tournamentId"]
	fmt.Print("ugabuga")
	fmt.Printf("%s %s", tournamentId, "id")

	var roundsNumber int
	queryRounds := fmt.Sprintf(`SELECT ROUNDS_NUMBER FROM TOURNAMENT WHERE TOURNAMENT_ID = '%s';`, tournamentId)

	err := db.QueryRow(queryRounds).Scan(&roundsNumber)
	if err != nil {
		log.Println("Ошибка при получения числа раундов:", err)
		http.Error(w, "Ошибка при получении информации о турнире", http.StatusInternalServerError)
		return
	}

	maxParticipants := 1 << roundsNumber

	var count int
	countq := fmt.Sprintf(`SELECT COUNT(*) FROM SELECTED_MEMBER WHERE TOURNAMENT_ID = '%s';`, tournamentId)

	err = db.QueryRow(countq).Scan(&count)
	if err != nil {
		log.Println("Ошибка при подсчете подтвержденных участников:", err)
		http.Error(w, "Ошибка при подсчете участников", http.StatusInternalServerError)
		return
	}

	if count < maxParticipants {
		http.Error(w, fmt.Sprintf("Недостаточное количество участников: ожидается %d, фактически %d", maxParticipants, count), http.StatusBadRequest)
		return
	}

	participants, err := getParticipants(tournamentId, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(1)
	for i := 1; i < len(participants); i++ {
		fmt.Print(i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(participants), func(i, j int) {
		participants[i], participants[j] = participants[j], participants[i]
	})

	for i, participant := range participants {
		matchNumber := (i / 2) + 1
		var column string
		if i%2 == 0 {
			column = "TOP_PARTICIPANT_ID"
		} else {
			column = "BOTTOM_PARTICIPANT_ID"
		}

		query := "UPDATE MATCH_TABLE SET " + column + " = '" + participant + "' WHERE MATCH_NUMBER = " + fmt.Sprint(matchNumber) + " AND TOURNAMENT_ID = '" + tournamentId + "'" // fmt.Print(query)
		fmt.Print(query)
		_, err := db.Exec(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func updateMatchScoreHandler(w http.ResponseWriter, r *http.Request) {
	var scoreUpdate MatchScoreUpdate
	err := json.NewDecoder(r.Body).Decode(&scoreUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	queryMatchExists := fmt.Sprintf("SELECT COUNT(*) FROM MATCH_TABLE WHERE MATCH_ID = '%s'", scoreUpdate.MatchID)
	var matchCount int
	err = db.QueryRow(queryMatchExists).Scan(&matchCount)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if matchCount == 0 {
		http.Error(w, "Матч не найден", http.StatusNotFound)
		return
	}

	var topParticipantID, bottomParticipantID string

	err = db.QueryRow(fmt.Sprintf("SELECT convert(nvarchar(50), TOP_PARTICIPANT_ID), convert(nvarchar(50),BOTTOM_PARTICIPANT_ID) FROM MATCH_TABLE WHERE MATCH_ID = '%s'", scoreUpdate.MatchID)).Scan(&topParticipantID, &bottomParticipantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	checkParticipantsQuery := fmt.Sprintf(`
                SELECT
                    COUNT(*)
                FROM
                    SELECTED_MEMBER
                WHERE
                    PARTICIPANT_ID IN ('%s', '%s')`, topParticipantID, bottomParticipantID)

	var participantCount int
	err = db.QueryRow(checkParticipantsQuery).Scan(&participantCount)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if participantCount < 2 {
		http.Error(w, "Один или оба бойца не найдены в списке участников.", http.StatusNotFound)
		return
	}

	var existingTopScore, existingBottomScore sql.NullInt64
	err = db.QueryRow("SELECT TOP_SCORE, BOTTOM_SCORE FROM MATCH_TABLE WHERE MATCH_ID = @p1", sql.Named("p1", scoreUpdate.MatchID)).Scan(&existingTopScore, &existingBottomScore)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if existingTopScore.Valid || existingBottomScore.Valid {
		http.Error(w, "Невозможно обновить счет: он уже задан", http.StatusConflict)
		return
	}

	query := fmt.Sprintf("UPDATE MATCH_TABLE SET TOP_SCORE = %d, BOTTOM_SCORE = %d WHERE MATCH_ID = '%s'",
		scoreUpdate.Top_score,
		scoreUpdate.Bottom_score,
		scoreUpdate.MatchID,
	)

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.QueryRow(fmt.Sprintf("SELECT convert(nvarchar(50), TOP_PARTICIPANT_ID), convert(nvarchar(50), BOTTOM_PARTICIPANT_ID) FROM MATCH_TABLE WHERE MATCH_ID = '%s'", scoreUpdate.MatchID)).Scan(&topParticipantID, &bottomParticipantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print("\n							", scoreUpdate)

	var winnerID string
	var loserID string
	if scoreUpdate.Top_score > scoreUpdate.Bottom_score {
		winnerID = topParticipantID
		loserID = bottomParticipantID
	} else {
		winnerID = bottomParticipantID
		loserID = topParticipantID
	}

	updateWinnerQuery := fmt.Sprintf("UPDATE MATCH_TABLE SET WINNER_ID = '%s' WHERE MATCH_ID = '%s'", winnerID, scoreUpdate.MatchID)
	_, err = db.Exec(updateWinnerQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Ошибка при обновлении WINNER_ID: %s", err.Error())
		return
	}

	q2 := `SELECT 
            MATCH_TABLE.ROUND_NUMBER, 
            TOURNAMENT.ROUNDS_NUMBER 
           FROM MATCH_TABLE 
           INNER JOIN TOURNAMENT 
           ON MATCH_TABLE.TOURNAMENT_ID = TOURNAMENT.TOURNAMENT_ID 
           WHERE MATCH_ID = @p1`
	var isLastMatch bool
	var r1, r2 int
	err = db.QueryRow(q2, sql.Named("p1", scoreUpdate.MatchID)).Scan(&r1, &r2)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r1 == r2 {
		isLastMatch = true
	} else {
		if r1 < r2 {
			isLastMatch = false
		} else {
			http.Error(w, "Что-то явно пошло не так с раундами", http.StatusInternalServerError)
			return
		}
	}

	if isLastMatch {
		insertWinnerQuery := fmt.Sprintf(`
            INSERT INTO RESULTS (TOURNAMENT_ID, ROUND_NUMBER, MATCH_ID, PARTICIPANT_ID, PLACE)
            SELECT mt.TOURNAMENT_ID, mt.ROUND_NUMBER, mt.MATCH_ID, '%s', 1
            FROM MATCH_TABLE mt
            WHERE mt.MATCH_ID = '%s'`, winnerID, scoreUpdate.MatchID)
		_, err = db.Exec(insertWinnerQuery)
		if err != nil {
			http.Error(w, "Ошибка при обновлении места победителя: "+err.Error(), http.StatusInternalServerError)
			return
		}

		insertLoserQuery := fmt.Sprintf(`
            INSERT INTO RESULTS (TOURNAMENT_ID, ROUND_NUMBER, MATCH_ID, PARTICIPANT_ID, PLACE)
            SELECT mt.TOURNAMENT_ID, mt.ROUND_NUMBER, mt.MATCH_ID, '%s', 2
            FROM MATCH_TABLE mt
            WHERE mt.MATCH_ID = '%s'`, loserID, scoreUpdate.MatchID)
		_, err = db.Exec(insertLoserQuery)
		if err != nil {
			http.Error(w, "Ошибка при обновлении места проигравшего: "+err.Error(), http.StatusInternalServerError)
			return
		}

		return
	} else {
		insertLoserQuery := fmt.Sprintf(`
            INSERT INTO RESULTS (TOURNAMENT_ID, ROUND_NUMBER, MATCH_ID, PARTICIPANT_ID, PLACE)
            SELECT mt.TOURNAMENT_ID, mt.ROUND_NUMBER, mt.MATCH_ID, '%s', 
            (SELECT COUNT(DISTINCT ROUND_NUMBER) 
             FROM MATCH_TABLE 
             WHERE TOURNAMENT_ID = mt.TOURNAMENT_ID) - mt.ROUND_NUMBER + 2
            FROM MATCH_TABLE mt
            WHERE mt.MATCH_ID = '%s'`, loserID, scoreUpdate.MatchID)
		_, err = db.Exec(insertLoserQuery)
		if err != nil {
			http.Error(w, "Ошибка при обновлении места проигравшего: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = assignNextMatch(winnerID, scoreUpdate.MatchID, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getParticipantsNamed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentId := vars["tournamentId"]
	rows, err := db.Query(`
        SELECT 
            CONVERT(NVARCHAR(100), SELECTED_MEMBER.PARTICIPANT_ID), 
            FIGHTER.NAME_TEXT 
        FROM 
            SELECTED_MEMBER 
        INNER JOIN 
            FIGHTER ON SELECTED_MEMBER.FIGHTER_ID = FIGHTER.FIGHTER_ID 
        WHERE 
            SELECTED_MEMBER.TOURNAMENT_ID = @p1`, sql.Named("p1", tournamentId))
	if err != nil {
		log.Print(err)
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var participants []Participant

	for rows.Next() {
		var participant Participant

		if err := rows.Scan(&participant.ParticipantID, &participant.Name); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}

		participants = append(participants, participant)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"participantlist": participants,
	}
	participantsJSON, err := json.MarshalIndent(response, "", "  ")
	log.Printf("%s", string(participantsJSON))

	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(participantsJSON)
}

func getMatchesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentID := vars["tournament_id"]
	query := fmt.Sprintf("SELECT CONVERT(NVARCHAR(100), MATCH_ID), CONVERT(NVARCHAR(100), TOURNAMENT_ID), ROUND_NUMBER, convert(nvarchar(50),TOP_PARTICIPANT_ID), TOP_SCORE, convert(nvarchar(50),BOTTOM_PARTICIPANT_ID), BOTTOM_SCORE, convert(nvarchar(50),WINNER_ID), WHEN_PLAYED, MATCH_NUMBER FROM MATCH_TABLE WHERE TOURNAMENT_ID = '%s';", tournamentID)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
		http.Error(w, "Ошибка при выполнении запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var matches []Match
	for rows.Next() {
		var match Match
		var topParticipantID, bottomParticipantID, winnerID sql.NullString
		var whenPlayed sql.NullString
		var topScore, bottomScore sql.NullInt64

		if err := rows.Scan(
			&match.MatchID,
			&match.TournamentID,
			&match.RoundNumber,
			&topParticipantID,
			&topScore,
			&bottomParticipantID,
			&bottomScore,
			&winnerID,
			&whenPlayed,
			&match.MatchNumber,
		); err != nil {
			http.Error(w, "Ошибка при сканировании результата: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if topParticipantID.Valid {
			match.TopParticipantID = topParticipantID.String
		}
		if topScore.Valid {
			match.Top_score = int(topScore.Int64)
		}
		if bottomParticipantID.Valid {
			match.BottomParticipantID = bottomParticipantID.String
		}
		if bottomScore.Valid {
			match.Bottom_score = int(bottomScore.Int64)
		}

		if winnerID.Valid {
			match.WinnerID = winnerID.String
		}

		if whenPlayed.Valid {
			match.WhenPlayed = whenPlayed.String
		}

		matches = append(matches, match)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при получении данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"matchlist": matches,
	}

	matchesJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s", string(matchesJSON))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(matchesJSON)
}

func transformPlace(place int, roundNumber int, maxRounds int) string {
	switch roundNumber {
	case maxRounds:
		if place == 1 {
			return "1"
		} else {
			if place == 2 {
				return "2"
			}
		}
	case maxRounds - 1:
		return "3-4"
	case maxRounds - 2:
		return "5-8"
	default:
		participants := int(math.Pow(2, float64(maxRounds-roundNumber)))
		startPlace := participants + 1
		endPlace := participants * 2
		return fmt.Sprintf("%d-%d", startPlace, endPlace)
	}
	return ""
}

func getResultsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentID := vars["tournament_id"]
	if tournamentID == "" {
		http.Error(w, "tournamentId параметр требуется", http.StatusBadRequest)
		return
	}

	query := `
        SELECT convert(nvarchar(50), TOURNAMENT_ID), ROUND_NUMBER, convert(nvarchar(50), MATCH_ID), convert(nvarchar(50), PARTICIPANT_ID), PLACE
        FROM RESULTS
        WHERE TOURNAMENT_ID = @p1
    `

	rows, err := db.Query(query, sql.Named("p1", tournamentID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при выполнении запроса: %s", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	query2 := "SELECT TOURNAMENT.ROUNDS_NUMBER FROM TOURNAMENT WHERE TOURNAMENT_ID = @p1"
	row, err := db.Query(query2, sql.Named("p1", tournamentID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при выполнении запроса: %s", err), http.StatusInternalServerError)
		return
	}
	defer row.Close()
	var maxRounds int
	for row.Next() {
		if err := row.Scan(&maxRounds); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обработке результата: %s", err), http.StatusInternalServerError)
			return
		}
	}

	var results []TableResults
	for rows.Next() {
		var result TableResults
		var tempPlace int
		if err := rows.Scan(&result.TournamentID, &result.RoundNumber, &result.MatchID, &result.ParticipantID, &tempPlace); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обработке результата: %s", err), http.StatusInternalServerError)
			return
		}

		result.Place = transformPlace(tempPlace, result.RoundNumber, maxRounds)
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при чтении строк результата: %s", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"results": results,
	}

	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func updateClub(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Connected")

	var club Club

	if err := json.NewDecoder(r.Body).Decode(&club); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании данных: %s", err), http.StatusBadRequest)
		return
	}

	if club.Title == "" || club.FoundedDate == "" || club.Rating == "" || club.Address == "" {
		http.Error(w, "Не все обязательные поля заполнены", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", club.FoundedDate); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	query := `
        UPDATE CLUB
        SET TITLE = @TITLE, DESCRIPTION_TEXT = @DESCRIPTION_TEXT, CONTACT = @CONTACT, RATING = @RATING, COUNTRY = @ADDRESS 
        WHERE CLUB_ID = @CLUB_ID`
	fmt.Print(club.ClubID)
	_, err := db.Exec(query, sql.Named("TITLE", club.Title),
		sql.Named("DESCRIPTION_TEXT", *club.DescriptionText),
		sql.Named("CONTACT", *club.Contact),
		sql.Named("RATING", club.Rating),
		sql.Named("ADDRESS", club.Address),
		sql.Named("CLUB_ID", club.ClubID))

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при обновлении клуба: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteClub(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "ID клуба не указан", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM CLUB WHERE CLUB_ID = @ID"

	_, err := db.Exec(query, sql.Named("ID", id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при удалении клуба: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func deleteFighter(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := "DELETE FROM FIGHTER WHERE FIGHTER_ID = @id"
	_, err := db.Exec(query, sql.Named("id", id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при удалении бойца: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateFighter(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Connected")

	var fighter Fighter

	if err := json.NewDecoder(r.Body).Decode(&fighter); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании данных: %s", err), http.StatusBadRequest)
		return
	}

	if fighter.Name == "" || fighter.Birthday == "" || fighter.Rating == "" || fighter.Country == "" {
		http.Error(w, "Не все обязательные поля заполнены", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", fighter.Birthday); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	query := `
        UPDATE FIGHTER
        SET NAME_TEXT = @NAME_TEXT, DESCRIPTION_TEXT = @DESCRIPTION_TEXT, CATEGORY = @CATEGORY, RATING = @RATING, COUNTRY = @COUNTRY, BIRTHDAY = @BIRTHDAY, CLUB_ID = @CLUB_ID 
        WHERE FIGHTER_ID = @FIGHTER_ID`

	fmt.Print(fighter.ClubID)
	_, err := db.Exec(query, sql.Named("NAME_TEXT", fighter.Name),
		sql.Named("DESCRIPTION_TEXT", *fighter.Description),
		sql.Named("CATEGORY", fighter.Category),
		sql.Named("RATING", fighter.Rating),
		sql.Named("COUNTRY", fighter.Country),
		sql.Named("BIRTHDAY", fighter.Birthday),
		sql.Named("CLUB_ID", fighter.ClubID),
		sql.Named("FIGHTER_ID", fighter.FighterID))

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при обновлении бойца: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateOrganazer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var organazer Organazer

	if err := json.NewDecoder(r.Body).Decode(&organazer); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании данных: %s", err), http.StatusBadRequest)
		return
	}

	if organazer.Title == "" || organazer.FoundationDate == "" || organazer.AddressText == "" {
		http.Error(w, "Не все обязательные поля заполнены", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", organazer.FoundationDate); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	query := `
        UPDATE ORGANAZER_TOUR
        SET TITLE = @TITLE, DESCRIPTION_TEXT = @DESCRIPTION_TEXT, FOUNDATION_DATE = @FOUNDATION_DATE, 
            ADDRESS_TEXT = @ADDRESS_TEXT, CONTACT_INFO = @CONTACT_INFO 
        WHERE ORGANAZER_ID = @ORGANAZER_ID`

	_, err := db.Exec(query, sql.Named("TITLE", organazer.Title),
		sql.Named("DESCRIPTION_TEXT", organazer.DescriptionText),
		sql.Named("FOUNDATION_DATE", organazer.FoundationDate),
		sql.Named("ADDRESS_TEXT", organazer.AddressText),
		sql.Named("CONTACT_INFO", organazer.ContactInfo),
		sql.Named("ORGANAZER_ID", id))

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при обновлении организатора: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteOrganazer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "ID организатора не указан", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM ORGANAZER_TOUR WHERE ORGANAZER_ID = @ID"

	_, err := db.Exec(query, sql.Named("ID", id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при удалении организатора: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteinviteRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	tournamentID := vars["tournamentId"]
	inviteId := vars["inviteId"]
	fmt.Print(tournamentID, "\n", inviteId)
	if tournamentID == "" || inviteId == "" {
		http.Error(w, "Missing tournament ID or invitedId", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM INVITED_PARTICIPANTS WHERE TOURNAMENT_ID = @p1 AND FIGHTER_ID = @p2"

	_, err := db.Exec(query, sql.Named("p1", tournamentID), sql.Named("p2", inviteId))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing query: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteTournament(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["tournamentId"]
	query := "DELETE FROM TOURNAMENT WHERE TOURNAMENT_ID = @p1"
	_, err := db.Exec(query, sql.Named("p1", id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при удалении турнира: %s", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateTournament(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["tournamentId"]
	var tournament Tournament

	if err := json.NewDecoder(r.Body).Decode(&tournament); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании данных: %s", err), http.StatusBadRequest)
		return
	}

	if tournament.Title == "" || tournament.StartDate == "" || tournament.EndDate == "" || tournament.Address == "" {
		http.Error(w, "Не все обязательные поля заполнены", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", tournament.StartDate); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	if _, err := time.Parse("2006-01-02", tournament.EndDate); err != nil {
		http.Error(w, "Неверный формат даты. Используйте формат YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	query := `
        UPDATE TOURNAMENT
        SET TITLE = @TITLE, START_DATE = @START_DATE, END_DATE = @END_DATE, ADDRESS = @ADDRESS WHERE TOURNAMENT_ID = @ID`

	_, err := db.Exec(query, sql.Named("TITLE", tournament.Title),
		sql.Named("START_DATE", tournament.StartDate),
		sql.Named("END_DATE", tournament.EndDate),
		sql.Named("ADDRESS", tournament.Address),
		sql.Named("ID", id))

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при обновлении турнира: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	requestId := mux.Vars(r)["requestId"]
	query := "DELETE FROM REQUEST WHERE REQUEST_id = @P1"
	_, err := db.Exec(query, sql.Named("p1", requestId))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при удалении заявки: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getRequestsByFighterId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fighterId := vars["id"]

	query := fmt.Sprintf(`
        SELECT 
            convert(nvarchar(50), r.REQUEST_ID), 
            convert(nvarchar(50), r.FIGHTER_ID), 
            t.TITLE, 
            convert(nvarchar(50), r.TOURNAMENT_ID), 
            t.CATEGORY, 
            r.DECISION 
        FROM REQUEST r
        JOIN TOURNAMENT t ON r.TOURNAMENT_ID = t.TOURNAMENT_ID
        WHERE r.FIGHTER_ID = '%s'`, fighterId)

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var requests []RequestWithTournamentTitle
	for rows.Next() {
		var req RequestWithTournamentTitle
		if err := rows.Scan(&req.RequestID, &req.FighterID, &req.Title, &req.TournamentID, &req.Category, &req.Decision); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requests = append(requests, req)
	}

	response := map[string]interface{}{
		"requests": requests,
	}

	requestsJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(requestsJSON)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user UserProfile

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании данных: %s", err), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Заполните все обязательные поля", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO USER_PROFILE (USERNAME, PROFILE_TYPE, E_MAIL, PROFILE_PASSWORD)
        VALUES (@USERNAME, @PROFILE_TYPE, @E_MAIL, @PROFILE_PASSWORD) 
    `

	_, err := db.Exec(query,
		sql.Named("USERNAME", user.Username),
		sql.Named("PROFILE_TYPE", user.ProfileType),
		sql.Named("E_MAIL", user.Email),
		sql.Named("PROFILE_PASSWORD", user.Password),
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при регистрации пользователя: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании данных: %s", err), http.StatusBadRequest)
		return
	}

	if request.Username == "" || request.Password == "" {
		http.Error(w, "Заполните все обязательные поля", http.StatusBadRequest)
		return
	}

	query := "SELECT USER_PROFILE_ID FROM USER_PROFILE WHERE USERNAME = @USERNAME AND PROFILE_PASSWORD = @PROFILE_PASSWORD"

	var userID string
	err := db.QueryRow(query, sql.Named("USERNAME", request.Username), sql.Named("PROFILE_PASSWORD", request.Password)).Scan(&userID)

	if err != nil {
		http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getUserProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := `
        SELECT 
            CONVERT(NVARCHAR(50), up.USER_PROFILE_ID),
            up.USERNAME,
            up.PROFILE_TYPE,
            up.E_MAIL,
            CONVERT(NVARCHAR(50), o.ORGANAZER_ID) AS Linked_id_1,
            CONVERT(NVARCHAR(50), c.CLUB_ID) AS Linked_id_2,
            CONVERT(NVARCHAR(50), f.FIGHTER_ID) AS Linked_id_3
        FROM USER_PROFILE up
        LEFT JOIN ORGANAZER_TOUR o ON up.USER_PROFILE_ID = o.USER_PROFILE_ID
        LEFT JOIN CLUB c ON up.USER_PROFILE_ID = c.USER_PROFILE_ID
        LEFT JOIN FIGHTER f ON up.USER_PROFILE_ID = f.USER_PROFILE_ID
        WHERE up.USER_PROFILE_ID = @USER_PROFILE_ID
    `

	var user UserProfile2
	var linkedID1, linkedID2, linkedID3 sql.NullString

	err := db.QueryRow(query, sql.Named("USER_PROFILE_ID", id)).
		Scan(&user.UserProfileID, &user.Username, &user.ProfileType, &user.Email,
			&linkedID1, &linkedID2, &linkedID3)

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при извлечении профиля: %s", err), http.StatusNotFound)
		return
	}

	user.Linked_id_1 = linkedID1.String
	if !linkedID1.Valid {
		user.Linked_id_1 = ""
	}

	user.Linked_id_2 = linkedID2.String
	if !linkedID2.Valid {
		user.Linked_id_2 = ""
	}

	user.Linked_id_3 = linkedID3.String
	if !linkedID3.Valid {
		user.Linked_id_3 = ""
	}

	response := map[string]interface{}{
		"user": user,
	}

	UserIDSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(UserIDSON)
}

func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user UserProfile
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при декодировании данных: %s", err), http.StatusBadRequest)
		return
	}
	query := ``
	if user.Password != "" {
		query = fmt.Sprintf(`UPDATE USER_PROFILE SET USERNAME = @USERNAME, PROFILE_TYPE = @PROFILE_TYPE, E_MAIL = '%s', PROFILE_PASSWORD = @PROFILE_PASSWORD WHERE USER_PROFILE_ID = @USER_PROFILE_ID`, user.Email)
		fmt.Print(query)
		_, err := db.Exec(query,
			sql.Named("USERNAME", user.Username),
			sql.Named("PROFILE_TYPE", user.ProfileType),
			sql.Named("PROFILE_PASSWORD", user.Password),
			sql.Named("USER_PROFILE_ID", id),
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обновлении профиля: %s", err), http.StatusInternalServerError)
			return
		}
	} else {
		query = fmt.Sprintf(`
        UPDATE USER_PROFILE SET USERNAME = @USERNAME, PROFILE_TYPE = @PROFILE_TYPE, E_MAIL = '%s' WHERE USER_PROFILE_ID = @USER_PROFILE_ID
    `, user.Email)
		_, err := db.Exec(query,
			sql.Named("USERNAME", user.Username),
			sql.Named("PROFILE_TYPE", user.ProfileType),
			sql.Named("USER_PROFILE_ID", id),
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при обновлении профиля: %s", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func deleteUserProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := "DELETE FROM USER_PROFILE WHERE USER_PROFILE_ID = @USER_PROFILE_ID"

	_, err := db.Exec(query, sql.Named("USER_PROFILE_ID", id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при удалении профиля: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["e_mail"]
	fmt.Print(email)
	query := fmt.Sprintf("SELECT CONVERT(NVARCHAR(50), USER_PROFILE_ID) FROM USER_PROFILE WHERE E_MAIL = '%s'", email)

	var user_id string
	err := db.QueryRow(query).Scan(&user_id)

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при извлечении профиля: %s", err), http.StatusNotFound)
		return
	}
	fmt.Print("ungabunga", user_id)
	response := map[string]interface{}{
		"user_id": user_id,
	}

	UserIDSON, err := json.MarshalIndent(response, "", "  ")
	fmt.Print("PO email", string(UserIDSON))
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(UserIDSON)
}

func getUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	query := "SELECT CONVERT(NVARCHAR(50), USER_PROFILE_ID) FROM USER_PROFILE WHERE USERNAME = @USERNAME"

	var user_id string
	err := db.QueryRow(query, sql.Named("USERNAME", username)).Scan(&user_id)

	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при извлечении профиля: %s", err), http.StatusNotFound)
		return
	}
	response := map[string]interface{}{
		"user_id": user_id,
	}

	UserIDSON, err := json.MarshalIndent(response, "", "  ")
	fmt.Print(string(UserIDSON))
	if err != nil {
		http.Error(w, "Ошибка при маршаллинге данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(UserIDSON)
}
