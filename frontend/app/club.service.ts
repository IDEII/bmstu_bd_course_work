import { Injectable } from '@angular/core';
import { Club } from './club'
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { ClubMember } from './club-member';
import { catchError, map, Observable, throwError } from 'rxjs';
import { Fighter } from './fighter';

@Injectable({
  providedIn: 'root'
})

export class ClubService {
  
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  getClubs(): Observable<Club[]> {
    return this.http.get(`${this.apiUrl}/clubs`).pipe(
      map((data: any) => {
        let clubsList = data["clublist"];
        return clubsList.map(function(club: any): Club {
          return new Club(
            club.user_id,
            club.id,
            club.title,
            club.description,
            club.address,
            club.contact,
            club.rating,
            new Date(club.founded_date)
          );
        });
      })
    );
  }
  getMembers(id: string): Observable<Fighter[]> {
    return this.http.get(`${this.apiUrl}/club-profile/${id}/members`).pipe(
      map((data: any) => {
        let memberlist = data["memberlist"];
        return memberlist.map(function(member: Fighter): Fighter {
          return new Fighter(
            member.user_id,
            member.club_id,
            member.id,
            member.name,
            member.description,
            member.country,
            new Date(member.birthday),
            member.rating,
            member.category,
          );
        });
      })
    );
  }
  getClub(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/club-profile/${id}`);
  }

  getFightersWithOutClub(): Observable<Fighter[]> {
    return this.http.get(`${this.apiUrl}/fighters_no_club`).pipe(
      map((data: any) => {
        let fighters = data["fighterlist"];
        return fighters.map(function(fighter: Fighter): Fighter {
          return new Fighter(
            fighter.user_id,
            fighter.club_id,
            fighter.id,
            fighter.name,
            fighter.description,
            fighter.country,
            new Date(fighter.birthday),
            fighter.rating,
            fighter.category,
          );
        });
      })
    );
  }

  addMemberToClub(id: string, fighter_id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/club-profile/${id}/addMember/${fighter_id}`);
  }

  kickFromClub(id:string, fighter_id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/club-profile/${id}/kickMember/${fighter_id}`);
  }

  updateClub(id: string, club: any): Observable<any> {
    const updateClubHeader = new HttpHeaders().set("Accept", "application/json");
    const data = JSON.stringify(club);
    return this.http.post(`${this.apiUrl}/club-profile/${id}/update`, data, { headers: updateClubHeader }).pipe(catchError(this.handleError)); // Обработка ошибок
}

  private handleError(error: any): Observable<never> {
    console.error('An error occurred:', error); // Log to console
    return throwError(() => new Error(error.message || 'Server error'));
  }

  addClub(club: any, user_id: string): Observable<any> {
    const newClub = new HttpHeaders().set("Accept", "application/json");
    const body = JSON.stringify(club)
    console.log(body)
    return this.http.post<any>(`${this.apiUrl}/add_club/${user_id}`, body, {headers: newClub}).pipe(catchError(this.handleError));
  }

  deleteClub(clubId: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/club-profile/${clubId}/delete`);
  }
}
