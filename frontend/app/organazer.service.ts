import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, Observable, throwError } from 'rxjs';
import { Organazer } from './organazer';
import { Tournament } from './tournament';

@Injectable({
  providedIn: 'root'
})
export class OrganazerService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  getOrganazers(): Observable<Organazer[]> {
    return this.http.get(`${this.apiUrl}/organazers`).pipe(
      map((data: any) => {
        let organazerlist = data["Organazerlist"];
        return organazerlist.map(function(organazer: any): Organazer {
          return new Organazer(
            organazer.user_id,
           organazer.id,
            organazer.title,
            organazer.description,
            organazer.address,
            organazer.contact,
            new Date(organazer.founded_date),
          );
        });
      })
    );
  }

  getOrganazer(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/organazer-profile/${id}`);
  }

  private handleError(error: any): Observable<never> {
    console.error('An error occurred:', error); 
    return throwError(() => new Error(error.message || 'Server error'));
  }

  addOrganazer(organazer: any, user_id: string): Observable<any> {
    const newOrganazer = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/add_organazer/${user_id}`, JSON.stringify(organazer), {headers: newOrganazer}).pipe(catchError(this.handleError));
  }
  
  getOrgTournaments(id: string): Observable<Tournament[]> {
    return this.http.get(`${this.apiUrl}/organazer-profile/${id}/tournaments_organazed`).pipe(
      map((data: any) => {
        let tournamentlistorg = data["tournamentlistorg"];
        return tournamentlistorg.map(function(tournament: any): Tournament {
          return new Tournament(
            tournament.id,
            tournament.name,
            tournament.address,
            new Date(tournament.startDate),
            new Date(tournament.endDate),
            tournament.organazer,
            tournament.rounds,
            tournament.category,
          );
        });
      })
    );
  }

  updateOrganazer(id: string, organazer: any): Observable<any> {
    const updateOrganazer = new HttpHeaders().set("Accept", "application/json");
    const data = JSON.stringify(organazer)
    return this.http.post(`${this.apiUrl}/organazer-profile/${id}/update`, data, {headers: updateOrganazer}).pipe(catchError(this.handleError));
  }

  deleteOrganazer(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/organazer-profile/${id}/delete`);
  }
}
