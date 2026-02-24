import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, Observable, of, throwError } from 'rxjs';
import { Tournament } from './tournament';
import { Request, InvitedParticipant, RequestWithTournamentTitle } from './request';

@Injectable({
  providedIn: 'root'
})
export class RequestManagerService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getFighterByRequestID(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/fighter-request/${id}`);
  }
 
  getRequests(tournamentId: string): Observable<Request[]> {
    const url = `${this.apiUrl}/requests/${tournamentId}`;
    return this.http.get(url).pipe(
      map((data: any) => {
        let requests = data["requests"];
        return requests.map(function(request: any): Request {
          return new Request(
            request.request_id,
            request.id_fighter,
            request.tournament_id,
            request.invited_id,
            request.decision
          );
        });
      }),
      catchError(this.handleError)
    );
  }

  getInvited(tournamentId: string): Observable<InvitedParticipant[]> {
    const url = `${this.apiUrl}/invited/${tournamentId}`;
    return this.http.get(url).pipe(
      map((data: any) => {
        let invitedParticipants = data["invitedParticipants"];
        return invitedParticipants.map(function(invitedParticipant: any): InvitedParticipant {
          return new InvitedParticipant(
            invitedParticipant.fighter_id,
            invitedParticipant.tournament_id,
            invitedParticipant.InvitedParticipant_id,
          );
        });
      }),
      catchError(this.handleError)
    );
  }

  approveRequest(requestId: string): Observable<any> {
    const body = {};  
    const headers = new HttpHeaders().set("Accept", "application/json");
    const url = `${this.apiUrl}/approve/${requestId}`;
    return this.http.post<any>(url, JSON.stringify(body), { headers: headers })
      .pipe(catchError(this.handleError));
  }

  denyRequest(requestId: string): Observable<any> {
    const body = {};  
    const headers = new HttpHeaders().set("Accept", "application/json");
    const url = `${this.apiUrl}/deny/${requestId}`;
    return this.http.post<any>(url, JSON.stringify(body), { headers: headers })
      .pipe(catchError(this.handleError));
  }

  private handleError(error: any): Observable<never> {
    console.error('An error occurred', error);
    return throwError(error.message || error);
  }

  getRequestsByFighterId(id: string): Observable<RequestWithTournamentTitle[]> {
    const url = `${this.apiUrl}/fighter/${id}/requests`;
    return this.http.get(url).pipe(
      map((data: any) => {
        let requests = data["requests"];
        return requests.map(function(request: any): RequestWithTournamentTitle {
          return new RequestWithTournamentTitle(
            request.request_id,
            request.id_fighter,
            request.name,
            request.tournament_id,
            request.category,
            request.decision
          );
        });
      }),
      catchError(this.handleError)
    );
  }
  deleteRequest(requestId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/requests/${requestId}/delete`);
  }
}