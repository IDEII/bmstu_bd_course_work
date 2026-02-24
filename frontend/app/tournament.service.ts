import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, Observable, of, throwError } from 'rxjs';
import { Match, Participant, TableResults, Tournament } from './tournament';

@Injectable({
  providedIn: 'root'
})
export class TournamentService {

  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  getTournaments(): Observable<Tournament[]> {
    return this.http.get(`${this.apiUrl}/tournaments`).pipe(
      map((data: any) => {
        let tournamentlist = data["tournamentlist"];
        return tournamentlist.map(function(tournament: any): Tournament {
          return new Tournament(
            tournament.id,
            tournament.name,
            tournament.address,
            new Date(tournament.startDate),
            new Date(tournament.endDate),
            tournament.organazer,
            tournament.rounds,
            tournament.category
          );
        });
      })
    );
  }
  getParticipants(tournamentId: string): Observable<Participant[]> {
    return this.http.get(`${this.apiUrl}/tournament/${tournamentId}/selected-members`).pipe(
      map((data: any) => {
        let participantlist = data["participantlist"];
        return participantlist.map(function(participant: any):Participant {
          return new Participant(
            participant.participantId,
            participant.name);});}));}

  getTournament(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/tournament/${id}`);
  }

  private handleError(error: any): Observable<never> {
    console.error('An error occurred:', error); 
    return throwError(() => new Error(error.message || 'Server error'));
  }

  addTournament(tournament: any): Observable<any> {
    const newTournament = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/organazer/create_tournament`, JSON.stringify(tournament), {headers: newTournament}).pipe(catchError(this.handleError));
  }
  
  getBracketData(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/tournament/${id}/bracket`);
  }

  sendRequest(fighterid: string, selectedTournamentid: any): Observable<any> {
    const body = {
      id_fighter: fighterid,
      tournament_id: selectedTournamentid
    };
    const newRequest = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/sendRequest`, JSON.stringify(body), {headers: newRequest}).pipe(catchError(this.handleError));
   }

   inviteRequest(tournamentid: string, fighterId: string): Observable<any> {
    const body = {
      fighterId: fighterId
    }
    const inviteRequest = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/tournament/${tournamentid}/invite`, JSON.stringify(body), {headers: inviteRequest}).pipe(catchError(this.handleError));
  }

  getSelectedMembers(id: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/tournament/${id}/selected-members`);
  }
  

  sendDrawRequest(tournamentId: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/tournament/${tournamentId}/draw`)
   }

   conductTournament(tournamentId: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/tournament/${tournamentId}/conduct`);
  }


  updateMatchScore(matchID: string, topScore: number, bottomScore: number, tournamentId:string): Observable<any> {
    const body = {
      match_id: matchID,
      top_score: topScore, 
      bottom_score: bottomScore
    };
    const updateMatchScore = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/tournament/${tournamentId}/update-match-score`, JSON.stringify(body), {headers: updateMatchScore}).pipe(catchError(this.handleError));
  }

  getMatchesById(tournament_id: string): Observable<Match[]> {
    return this.http.get(`${this.apiUrl}/tournament/${tournament_id}/matches`).pipe(
      map((data: any) => {
        let matchlist = data["matchlist"];
        return matchlist.map(function(match: any): Match {
          return new Match(
            match.match_id,
            match.tournament_id,
            match.round_number,
            match.match_number,
            match.top_participant_id,
            match.top_score,
            match.bottom_participant_id,
            match.bottom_score,
            match.winner_id,
            new Date(match.when_played)
          );
        });
      })
    );  
  }
         
  
  fetchResults(tournament_id: string): Observable<TableResults[]> {
    return this.http.get<any>(`${this.apiUrl}/tournament/${tournament_id}/results`).pipe(
      map((data: any) => {
        let resultsList = data["results"];
        return resultsList.map((result: any): TableResults => {
          return new TableResults(
            result.tournament_id,
            result.roundNumber,
            result.match_id,
            result.participant_id,
            result.place
          );
        });
      }),
      catchError(error => {
        console.error('Ошибка при получении данных:', error);
        return throwError('Ошибка при получении данных');
      })
    );
  }    

  updateSelectedMembers(tournamentId: string, members: any[]): Observable<any> {
    return this.http.put(`${this.apiUrl}/update_selected_members/${tournamentId}`, members).pipe(
      catchError(error => {
        console.error('Ошибка при обновлении участников:', error);
        throw error;
      })
    );
  }

  updateTournament(id: string, tournament: any): Observable<any> {
    const updateTournamentHeader = new HttpHeaders().set("Accept", "application/json");
    const body = JSON.stringify(tournament);
    return this.http.post(`${this.apiUrl}/tournament/${id}/update`, body, {headers: updateTournamentHeader}).pipe(
      catchError(error => {
        console.error('Ошибка при обновлении турнира:', error);
        throw error;
      })
    );
  }  
  deleteTournament(id: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/tournament/${id}/delete`).pipe();
  }

deleteInvite(tournament_id: string, id: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/tournament/${tournament_id}/delete_invite/${id}`);
  }

  
} 

